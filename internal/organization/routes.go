package organization

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Group, h *Handler) {
	e.GET("/organization", h.GetAll)
	e.GET("/organization/:category", h.GetByCategory)
	e.POST("/organization", h.Create)       // opsional, admin only
	e.PUT("/organization/:id", h.Update)    // opsional
	e.DELETE("/organization/:id", h.Delete) // opsional
}
