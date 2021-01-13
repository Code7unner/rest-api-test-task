package handlers

import (
	"github.com/code7unner/rest-api-test-task/todos/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TodoRequest struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	TimeToComplete string `json:"time_to_complete"`
}

type TodosHandler struct {
	service service.Service
}

func NewTodosHandler(s service.Service) *TodosHandler {
	return &TodosHandler{service: s}
}

// CreateTodo godoc
// @Summary Create new todo task
// @Description Create new todo task for current user
// @Tags todos
// @ID create-todo
// @Produce json
// @Param request body TodoRequest true "Request body"
// @Success 201
// @Security Bearer
// @Router /todo [post]
func (h TodosHandler) CreateTodo(c echo.Context) error {
	todoRequest := new(TodoRequest)
	if err := c.Bind(todoRequest); err != nil {
		return err
	}

	userID, err := getUserIDFromToken(c.Get("user"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	timeToComplete, err := parseDateTime(todoRequest.TimeToComplete)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("wrong event date"))
	}

	_, err = h.service.CreateTodo(
		userID,
		todoRequest.Title,
		todoRequest.Description,
		timeToComplete,
	)

	switch err {
	case service.ErrTodoNotCreated:
		return c.JSON(http.StatusBadRequest, err)
	case nil:
		return c.JSON(http.StatusCreated, nil)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}

// UpdateTodo godoc
// @Summary Update todo task
// @Description Update todo task for current user
// @Tags todos
// @ID update-todo
// @Param id path int true "Todo ID"
// @Success 204
// @Security Bearer
// @Router /todo/{id} [patch]
func (h TodosHandler) UpdateTodo(c echo.Context) error {
	todoRequest := new(TodoRequest)
	if err := c.Bind(todoRequest); err != nil {
		return err
	}

	userID, err := getUserIDFromToken(c.Get("user"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("wrong todo id"))
	}

	timeToComplete, err := parseDateTime(todoRequest.TimeToComplete)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("wrong event date"))
	}

	_, err = h.service.UpdateTodo(
		todoID,
		userID,
		todoRequest.Title,
		todoRequest.Description,
		timeToComplete,
	)

	switch err {
	case service.ErrTodoNotFound:
		return c.JSON(http.StatusBadRequest, err)
	case service.ErrTodoNotUpdated:
		return c.JSON(http.StatusBadRequest, err)
	case nil:
		return c.JSON(http.StatusNoContent, nil)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}

// DeleteTodo godoc
// @Summary Delete todo task
// @Description Delete todo task for current user
// @Tags todos
// @ID delete-todo
// @Param id path int true "Todo ID"
// @Produce json
// @Success 204
// @Security Bearer
// @Router /todo/:id [delete]
func (h TodosHandler) DeleteTodo(c echo.Context) error {
	todoRequest := new(TodoRequest)
	if err := c.Bind(todoRequest); err != nil {
		return err
	}

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("wrong todo id"))
	}

	err = h.service.DeleteTodo(todoID)
	switch err {
	case service.ErrTodoNotFound:
		return c.JSON(http.StatusBadRequest, err)
	case service.ErrTodoNotDeleted:
		return c.JSON(http.StatusBadRequest, err)
	case nil:
		return c.JSON(http.StatusNoContent, nil)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}

// GetAllTodos godoc
// @Summary Get all todos
// @Description Gets all todos for current user
// @Tags todos
// @ID get-all-todos
// @Produce json
// @Success 200 {object} []models.Todos
// @Security Bearer
// @Router /todo/all [get]
func (h TodosHandler) GetAllTodos(c echo.Context) error {
	userID, err := getUserIDFromToken(c.Get("user"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	todos, err := h.service.GetAllTodos(userID)
	switch err {
	case service.ErrTodosNotFound:
		return c.JSON(http.StatusBadRequest, err)
	case nil:
		return c.JSON(http.StatusOK, todos)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}

type CurrentTodoRequest struct {
	Time string `json:"time"`
}

// GetAllCurrentTodos godoc
// @Summary Get all current todos
// @Description Gets all current todos for current user
// @Tags todos
// @ID get-all-current-todos
// @Produce json
// @Param request body CurrentTodoRequest true "Request body"
// @Success 200 {object} []models.Todos
// @Security Bearer
// @Router /todo/current [post]
func (h TodosHandler) GetAllCurrentTodos(c echo.Context) error {
	currentTodoRequest := new(CurrentTodoRequest)
	if err := c.Bind(currentTodoRequest); err != nil {
		return err
	}

	userID, err := getUserIDFromToken(c.Get("user"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	timeToComplete, err := parseDateTime(currentTodoRequest.Time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("wrong event date"))
	}

	todos, err := h.service.GetAllCurrentTodos(userID, timeToComplete)
	switch err {
	case service.ErrTodosNotFound:
		return c.JSON(http.StatusBadRequest, err)
	case nil:
		return c.JSON(http.StatusOK, todos)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}
