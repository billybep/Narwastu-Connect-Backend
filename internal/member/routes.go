package member

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Group, h *Handler) {
	e.GET("/birthday/weekly", h.GetWeeklyBirthdays)
	e.POST("/members", h.CreateMember)
}
