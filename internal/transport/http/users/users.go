package users

import (
	"api/internal/application"
	"api/internal/transport/http/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase application.IUsers
}

func Register(g *echo.Group, useCase application.IUsers, middlewares middlewares.IMiddlewares) {
	h := &handler{useCase: useCase}
	router := g.Group("/users")

	router.GET("/:id", h.getOne)

	{
		_ = router.Group("", middlewares.Authenticate)
	}
}

func (h *handler) getOne(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(400)
	}

	user, err := h.useCase.GetOneById(userId)
	if err != nil {
		return c.NoContent(400)
	}

	return c.JSON(200, user)
}
