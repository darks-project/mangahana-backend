package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type handler struct {
	router *echo.Echo
}

func New() *handler {
	return &handler{
		router: echo.New(),
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
	baseUrl := h.router.Group("/api/v1")
	h.routesRegister(baseUrl)

	return h.router
}

func (h *handler) routesRegister(r *echo.Group) {

}
