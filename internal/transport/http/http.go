package http

import (
	"api/internal/application"
	"api/internal/transport/http/users"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase *application.UseCase
	router  *echo.Echo
}

func New(useCase *application.UseCase) *handler {
	return &handler{
		useCase: useCase,
		router:  echo.New(),
	}
}

func (h *handler) Run(socket string) error {
	server := &http.Server{
		Addr:         socket,
		Handler:      h.getRouter(),
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}
	return server.ListenAndServe()
}

func (h *handler) getRouter() http.Handler {
	h.routesRegister()

	return h.router
}

func (h *handler) routesRegister() {
	api := h.router.Group("/api/v1")

	users.Register(api, h.useCase.Users, h)
}
