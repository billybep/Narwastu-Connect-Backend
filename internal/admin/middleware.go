package admin

import (
	"net/http"
	"strings"

	"app/internal/auth"

	"github.com/labstack/echo/v4"
)

// ✅ Middleware untuk validasi JWT khusus Admin Dashboard
func AdminJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing token"})
		}

		tokenStr := extractToken(authHeader)
		claims, err := auth.ParseAdminToken(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid or expired token"})
		}

		// Simpan informasi admin ke context agar bisa digunakan handler selanjutnya
		c.Set("admin_id", claims.AdminID)
		c.Set("admin_email", claims.Email)
		c.Set("admin_role", claims.Role)

		return next(c)
	}
}

// ✅ Middleware tambahan untuk cek role tertentu (misalnya hanya System Admin)
func RequireAdminRole(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			roleVal := c.Get("admin_role")
			if roleVal == nil {
				return c.JSON(http.StatusForbidden, echo.Map{"error": "role not found"})
			}

			role := roleVal.(string)
			for _, allowed := range roles {
				if role == allowed {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, echo.Map{"error": "access denied"})
		}
	}
}

// helper untuk ambil token dari "Bearer <token>"
func extractToken(header string) string {
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimPrefix(header, "Bearer ")
	}
	return header
}
