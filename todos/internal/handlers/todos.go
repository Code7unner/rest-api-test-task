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
	case nil:
		return c.JSON(http.StatusCreated, nil)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}

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
	case nil:
		return c.JSON(http.StatusCreated, nil)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}
