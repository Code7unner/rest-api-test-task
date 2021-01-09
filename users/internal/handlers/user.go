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

func (h UserHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.service.GetUser(id)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, user)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}