package server

import (
	_ "github.com/code7unner/rest-api-test-task/todos/docs"
	"github.com/code7unner/rest-api-test-task/todos/internal/handlers"
	"github.com/code7unner/rest-api-test-task/todos/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// New returns new echo server.
func New(s service.Service, jwtSecret []byte) *echo.Echo {
	e := echo.New()

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}}))

	JWTmiddleware := middleware.JWT(jwtSecret)

	ht := handlers.NewTodosHandler(s)
	t := e.Group("/todo")
	t.Use(JWTmiddleware)
	t.POST("", ht.CreateTodo)
	t.PATCH("/:id", ht.UpdateTodo)
	t.DELETE("/:id", ht.DeleteTodo)
	t.GET("/all", ht.GetAllTodos)
	t.POST("/current", ht.GetAllCurrentTodos)

	return e
}
