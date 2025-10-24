package admin

import (
	admin "app/internal/admin"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, h *Handler) {
	protected := g.Group("/dashboard", admin.AdminJWTMiddleware)
	protected.GET("", h.GetDashboard)
}
