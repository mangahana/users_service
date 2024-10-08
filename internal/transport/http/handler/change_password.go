package handler

import (
	"users_service/internal/core"
	"users_service/internal/models"

	"github.com/labstack/echo/v4"
)

func (h *Handler) changePassword(ctx echo.Context) error {
	var dto *core.ChangePasswordDTO

	if err := ctx.Bind(&dto); err != nil {
		return ctx.NoContent(400)
	}

	if err := h.validator.Struct(dto); err != nil {
		return ctx.NoContent(400)
	}

	session := ctx.Get("user").(models.Session)

	err := h.useCase.ChangePassword(ctx.Request().Context(), session.UserID, dto)
	if err != nil {
		return ctx.NoContent(400)
	}

	return ctx.NoContent(200)
}
