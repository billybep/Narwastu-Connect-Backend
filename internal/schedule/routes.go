package schedule

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Group, h *Handler) {
	e.GET("/schedules", h.GetAll)
	e.GET("/schedules/latest", h.GetLatest)
	e.GET("/schedules/:date", h.GetByDate) // format: YYYY-MM-DD
	e.POST("/schedules", h.Create)
}
