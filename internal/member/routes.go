package member

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Group, h *Handler) {
	e.GET("/birthday/weekly", h.GetWeeklyBirthdays)
	e.GET("/members", h.GetMembers)
	e.GET("/members/:id", h.GetMemberByID)
	e.PUT("/members/:id", h.UpdateMember)
	e.POST("/members", h.CreateMember)
	e.POST("/members/:id/avatar", h.UploadAvatar)
	e.DELETE("/members/:id", h.DeleteMember)
}
