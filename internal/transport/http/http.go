package http

import (
	"context"
	"users_service/internal/application"
	"users_service/internal/transport/http/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	server       *echo.Echo
	useCase      application.UseCase
	uploadFolder string
}

func New(useCase application.UseCase, uploadFolder string) *HttpServer {
	return &HttpServer{
		server:       echo.New(),
		useCase:      useCase,
		uploadFolder: uploadFolder,
	}
}

func (h *HttpServer) ListenAndServe(socket string) {
	h.server.Start(socket)
}

func (h *HttpServer) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

func (h *HttpServer) Register() {
	h.server.Use(middleware.BodyLimit("10M"))
	baseUrl := h.server.Group("/api/v1/users")

	handler := handler.New(baseUrl, h.useCase, h.uploadFolder)
	handler.Register()
}
