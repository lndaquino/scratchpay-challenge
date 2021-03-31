package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
)

func EnsureIsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearerToken := c.Request().Header.Get("Authorization")
		token := strings.Split(bearerToken, " ")
		if len(token) != 2 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Provide the authorization bearer token",
			})
		}
		if token[0] != "Bearer" || token[1] != os.Getenv("APP_SECRET") {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Provide the authorization token",
			})
		}

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
