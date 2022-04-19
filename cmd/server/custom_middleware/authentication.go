package custom_middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/pkg/api"
)

func UserCheck(api api.UserAPI) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqToken := c.Request().Header.Get("Authorization")

			if reqToken == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Token required",
				})
			}

			splitToken := strings.Split(reqToken, "Bearer ")

			if len(splitToken) < 1 {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid token",
				})
			}

			reqToken = splitToken[1]
			user, ok := api.CheckTokenAndGetUser(reqToken)

			if ok && user.Username != "" {
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Token expired",
			})
		}
	}
}
