package common

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// StrictJWTMiddleware → hanya untuk protected routes
func StrictJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authz := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authz, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing bearer token"})
		}
		tokenStr := strings.TrimPrefix(authz, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("claims", claims)
		}
		return next(c)
	}
}

// OptionalJWTMiddleware → untuk public routes, tidak wajib token
func OptionalJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authz := c.Request().Header.Get("Authorization")
		if strings.HasPrefix(authz, "Bearer ") {
			tokenStr := strings.TrimPrefix(authz, "Bearer ")

			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})
			if err == nil && token.Valid {
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					c.Set("claims", claims)
				}
			}
			// kalau error → abaikan, route tetap jalan
		}
		return next(c)
	}
}
