package handlers

import (
	"github.com/code7unner/rest-api-test-task/users/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service service.Service
}

func NewUserHandler(s service.Service) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetCurrentUser(c echo.Context) error {
	userID, err := getUserIDFromToken(c.Get("user"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := h.service.GetUser(userID)
	switch err {
	case service.ErrUserNotFound:
		return c.JSON(http.StatusNotFound, nil)
	case nil:
		return c.JSON(http.StatusOK, user)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}

func (h UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("wrong ID"))
	}

	user, err := h.service.GetUser(id)
	switch err {
	case service.ErrUserNotFound:
		return c.JSON(http.StatusNotFound, nil)
	case nil:
		return c.JSON(http.StatusOK, user)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}
