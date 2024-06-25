package middleware

import (
	"github.com/labstack/echo/v4"
)

func RecoveryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//lo que escribis aca pasa antes de ejecutar
		return next(c)
		//lo que pasa aca despues
	}
}
