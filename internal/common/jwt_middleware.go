package common

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authz := c.Request().Header.Get("Authorization")
		if authz == "" {
			// Tidak ada header Authorization -> langsung reject
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing bearer token"})
		}
		if !strings.HasPrefix(authz, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization header"})
		}

		tokenStr := strings.TrimPrefix(authz, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid claims"})
		}

		c.Set("claims", claims)
		return next(c)
	}
}
