package middlewares

import "github.com/labstack/echo/v4"

type IMiddlewares interface {
	Authenticate(next echo.HandlerFunc) echo.HandlerFunc
}
