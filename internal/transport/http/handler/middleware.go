package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		session, err := h.useCase.GetSession(c.Request().Context(), token)
		if err != nil {
			return c.String(403, "FORBIDDEN")
		}

		c.Set("user", session)

		return next(c)
	}
}
