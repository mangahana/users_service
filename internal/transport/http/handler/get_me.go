package handler

import (
	"users_service/internal/models"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getMe(ctx echo.Context) error {
	session := ctx.Get("user").(models.Session)

	data, err := h.useCase.GetOne(ctx.Request().Context(), session.UserID)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.JSON(200, data)
}
