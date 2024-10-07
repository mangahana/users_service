package handler

import (
	"log"
	"users_service/internal/core"

	"github.com/labstack/echo/v4"
)

func (h *Handler) confirmation(ctx echo.Context) error {
	var dto core.ConfirmationDTO
	if err := ctx.Bind(&dto); err != nil {
		log.Panicln(err)
		return ctx.NoContent(400)
	}

	if err := h.validator.Struct(dto); err != nil {
		log.Panicln(err)
		return ctx.NoContent(400)
	}

	if err := h.useCase.Confirmation(ctx.Request().Context(), &dto); err != nil {
		log.Panicln(err)
		return ctx.NoContent(400)
	}

	return ctx.NoContent(200)
}
