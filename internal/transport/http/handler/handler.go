package handler

import (
	"users_service/internal/application"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	validator *validator.Validate
	router    *echo.Group
	useCase   application.UseCase
}

func New(router *echo.Group, useCase application.UseCase) *Handler {
	return &Handler{
		validator: validator.New(),
		router:    router,
		useCase:   useCase,
	}
}

func (h *Handler) Register() {
	h.router.GET("/phoneExists", h.phoneExists)
	h.router.POST("/join", h.join)
	h.router.POST("/confirmation", h.confirmation)
	h.router.POST("/register", h.register)
	h.router.POST("/signin", h.signIn)
}
