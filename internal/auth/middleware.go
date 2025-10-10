package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AdminJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing or invalid token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ParseAdminToken(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
		}

		// Simpan claims ke context untuk dipakai handler selanjutnya
		c.Set("admin_id", claims.AdminID)
		c.Set("admin_role", claims.Role)
		c.Set("admin_email", claims.Email)

		return next(c)
	}
}

// Role check middleware (e.g. hanya system_administrator)
func RequireRole(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := c.Get("admin_role")
			if role == nil {
				return c.JSON(http.StatusForbidden, echo.Map{"error": "no role in context"})
			}

			for _, allowed := range roles {
				if role == allowed {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, echo.Map{"error": "access denied"})
		}
	}
}
