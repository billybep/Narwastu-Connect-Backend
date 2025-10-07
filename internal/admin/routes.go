package admin

import "github.com/labstack/echo/v4"

func RegisterRoutes(g *echo.Group, h *Handler) {
	g.POST("/login", h.Login)

	protected := g.Group("/manage") // nanti tambahkan middleware role
	protected.GET("/admins", h.GetAll)
}
