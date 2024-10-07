package handler

import (
	"users_service/internal/core"

	"github.com/labstack/echo/v4"
)

func (h *Handler) signIn(ctx echo.Context) error {
	var dto core.SignInDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.NoContent(400)
	}

	if err := h.validator.Struct(dto); err != nil {
		return ctx.NoContent(400)
	}

	token, err := h.useCase.SignIn(ctx.Request().Context(), &dto)
	if err != nil {
		return ctx.String(400, err.Error())
	}

	return ctx.JSON(200, map[string]any{"token": token})
}
