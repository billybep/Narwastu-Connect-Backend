package admin

import (
	adminmiddleware "app/internal/admin"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, h *Handler) {
	// Group untuk Schedule Reports (read-only, bisa diakses semua role)
	reportGroup := g.Group(
		"/list/schedules",
		adminmiddleware.AdminJWTMiddleware, // hanya butuh JWT
	)
	reportGroup.GET("", h.GetSchedules)
	reportGroup.GET("/:id", h.GetScheduleByID)

	// Group khusus admin (untuk create)
	adminGroup := g.Group(
		"/schedules",
		adminmiddleware.AdminJWTMiddleware,
		adminmiddleware.RequireAdminRole("system_administrator", "administrator"),
	)
	adminGroup.POST("", h.CreateSchedule)
	adminGroup.PUT("/:id", h.UpdateSchedule)
	adminGroup.DELETE("/:id", h.DeleteSchedule)
}
