package handler

import (
	"users_service/internal/core"
	"users_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *Handler) join(ctx echo.Context) error {
	var dto core.JoinDTO

	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(400, cerror.InvalidData().Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		return ctx.JSON(400, cerror.InvalidData().Error())
	}

	err := h.useCase.Join(ctx.Request().Context(), dto.Phone)
	if err != nil {
		return ctx.JSON(400, err.Error())
	} else {
		return ctx.NoContent(200)
	}
}
