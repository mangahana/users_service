package handler

import (
	"users_service/internal/core"

	"github.com/labstack/echo/v4"
)

func (h *Handler) join(ctx echo.Context) error {
	var dto core.JoinDTO

	if err := ctx.Bind(&dto); err != nil {
		return ctx.NoContent(400)
	}

	if err := h.validator.Struct(dto); err != nil {
		return ctx.NoContent(400)
	}

	err := h.useCase.Join(ctx.Request().Context(), dto.Phone)

	if err != nil {
		return ctx.NoContent(400)
	} else {
		return ctx.NoContent(200)
	}
}
