package server

import (
	"github.com/code7unner/rest-api-test-task/users/internal/handlers"
	"github.com/code7unner/rest-api-test-task/users/internal/service"
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

	ha := handlers.NewAuthHandler(s)
	a := e.Group("/auth")
	a.POST("/register", ha.Register)
	a.POST("/login", ha.Login)

	hu := handlers.NewUserHandler(s)
	u := e.Group("/user")
	u.Use(JWTmiddleware)
	u.GET("/:id", hu.GetUser)

	return e
}
