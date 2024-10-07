package http

import (
	"context"
	"users_service/internal/application"
	"users_service/internal/transport/http/handler"

	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	server  *echo.Echo
	useCase application.UseCase
}

func New(useCase application.UseCase) *HttpServer {
	return &HttpServer{
		server:  echo.New(),
		useCase: useCase,
	}
}

func (h *HttpServer) ListenAndServe(socket string) {
	h.server.Start(socket)
}

func (h *HttpServer) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

func (h *HttpServer) Register() {
	baseUrl := h.server.Group("/api/v1")

	handler := handler.New(baseUrl, h.useCase)
	handler.Register()
}
