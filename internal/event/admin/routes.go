package admin

import (
	adminmiddleware "app/internal/admin"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, h *Handler) {
	// All admin
	eventAdminGroup := g.Group("/list", adminmiddleware.AdminJWTMiddleware)
	eventAdminGroup.GET("/events", h.GetAllEvents)
	eventAdminGroup.GET("/events/:id", h.GetEventByID)

	// Protected Admin Routes
	adminGroup := g.Group(
		"/events",
		adminmiddleware.AdminJWTMiddleware,
		adminmiddleware.RequireAdminRole("system_administrator", "administrator"),
	)

	adminGroup.POST("", h.CreateEvent)
	adminGroup.PUT("/:id", h.UpdateEvent)
	adminGroup.DELETE("/:id", h.DeleteEvent)
}
