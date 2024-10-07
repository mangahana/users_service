package handler

import "github.com/labstack/echo/v4"

func (h *Handler) phoneExists(ctx echo.Context) error {
	phone := ctx.QueryParam("phone")
	if err := h.validator.Var(phone, "required,min=10,max=10"); err != nil {
		return ctx.NoContent(400)
	}

	exists, err := h.useCase.PhoneExists(ctx.Request().Context(), phone)
	if err != nil {
		return ctx.NoContent(500)
	}

	return ctx.JSON(200, exists)
}
