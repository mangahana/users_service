package handler

import (
	"io"
	"net/http"
	"slices"
	"users_service/internal/core"
	"users_service/internal/core/cerror"
	"users_service/internal/models"

	"github.com/labstack/echo/v4"
)

var allowedMimeTypes = []string{"image/jpeg", "image/png", "image/webp"}

var exts = map[string]string{
	"image/jpeg": ".jpeg",
	"image/png":  ".png",
	"image/webp": ".webp",
}

func (h *Handler) uploadPhoto(ctx echo.Context) error {
	data, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(400, cerror.InvalidData().Error())
	}

	mime := http.DetectContentType(data)

	if !slices.Contains(allowedMimeTypes, mime) {
		return ctx.JSON(400, cerror.UNSUPPORTED_FORMAT)
	}

	randomStr, err := core.GenerateRandomString(64)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	filename := randomStr + exts[mime]

	user := ctx.Get("user").(models.Session)

	err = h.useCase.UploadPhoto(ctx.Request().Context(), user.UserID, filename, mime, data)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.NoContent(200)
}
