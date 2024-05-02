package http

import "github.com/labstack/echo/v4"

func (h *handler) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		user, err := h.useCase.Users.GetBySession(token)
		if err != nil {
			return c.String(401, "Unauthorized")
		}

		c.Set("userId", user.Id)

		return next(c)
	}
}
