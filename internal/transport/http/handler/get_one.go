package handler

import (
	"strconv"
	"users_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getOne(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(400, cerror.InvalidData().Error())
	}

	user, err := h.useCase.GetOne(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.JSON(200, user)
}
