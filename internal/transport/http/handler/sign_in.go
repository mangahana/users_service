package handler

import (
	"users_service/internal/core"
	"users_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *Handler) signIn(ctx echo.Context) error {
	var dto core.SignInDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(400, cerror.InvalidData().Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		return ctx.JSON(400, cerror.InvalidData().Error())
	}

	token, err := h.useCase.SignIn(ctx.Request().Context(), &dto)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.JSON(200, map[string]any{"token": token})
}
