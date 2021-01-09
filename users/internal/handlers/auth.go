package handlers

import (
	"github.com/code7unner/rest-api-test-task/users/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	service service.Service
}

func NewAuthHandler(s service.Service) *AuthHandler {
	return &AuthHandler{service: s}
}

// Register godoc
// @Summary Returns status code
// @Description register user in db
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Request body"
// @Success 201
// @Router /register [post]
func (h AuthHandler) Register(c echo.Context) error {
	request := new(RegisterRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := h.service.Register(request.Username, request.Password)
	switch err {
	case service.ErrUserCreating:
		return c.JSON(http.StatusBadRequest, nil)
	case service.ErrUserPasswordInvalid:
		return c.JSON(http.StatusBadRequest, nil)
	case nil:
		return c.JSON(http.StatusCreated, nil)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}

// Login godoc
// @Summary Returns access token
// @Description get token for user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Login body"
// @Success 200
// @Router /login [post]
func (h AuthHandler) Login(c echo.Context) error {
	request := new(RegisterRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := h.service.Login(request.Username, request.Password)
	switch err {
	case service.ErrUserNotFound:
		return c.JSON(http.StatusNotFound, nil)
	case service.ErrUserPasswordInvalid:
		return c.JSON(http.StatusBadRequest, nil)
	case service.ErrUserCreateJWTToken:
		return c.JSON(http.StatusBadRequest, nil)
	case nil:
		return c.JSON(http.StatusOK, token)
	default:
		return c.JSON(http.StatusInternalServerError, errorResponse("unexpected error"))
	}
}
