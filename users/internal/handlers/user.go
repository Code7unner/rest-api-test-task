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

// GetCurrentUser godoc
// @Summary Returns current user
// @Description Get current user
// @Tags users
// @ID get-current-user
// @Accept json
// @Produce json
// @Success 200 {object} models.Users
// @Security Bearer
// @Router /user [get]
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

// GetUser godoc
// @Summary Returns user
// @Description Get user
// @Tags users
// @ID get-user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.Users
// @Security Bearer
// @Router /user/{id} [get]
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
