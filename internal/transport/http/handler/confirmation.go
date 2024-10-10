package handler

import (
	"users_service/internal/core"
	"users_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *Handler) confirmation(ctx echo.Context) error {
	var dto core.ConfirmationDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(400, cerror.InvalidData().Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		return ctx.JSON(400, cerror.InvalidData().Error())
	}

	if err := h.useCase.Confirmation(ctx.Request().Context(), &dto); err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.NoContent(200)
}
