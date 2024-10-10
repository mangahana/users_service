package handler

import (
	"users_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *Handler) phoneExists(ctx echo.Context) error {
	phone := ctx.QueryParam("phone")
	if err := h.validator.Var(phone, "required,min=10,max=10"); err != nil {
		return ctx.JSON(400, cerror.InvalidData().Error())
	}

	exists, err := h.useCase.PhoneExists(ctx.Request().Context(), phone)
	if err != nil {
		return ctx.JSON(500, cerror.InvalidData().Error())
	}

	return ctx.JSON(200, exists)
}
