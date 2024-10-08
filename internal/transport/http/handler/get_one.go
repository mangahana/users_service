package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getOne(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.NoContent(400)
	}

	user, err := h.useCase.GetOne(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.NoContent(404)
	}

	return ctx.JSON(200, user)
}
