package admin

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, h *Handler) {
	// Public (tanpa token)
	g.POST("/login", h.Login)

	// Protected with token
	protected := g.Group("/manage", AdminJWTMiddleware)
	protected.POST("/create", h.CreateAdmin)
	protected.GET("/admins", h.GetAll)
}
