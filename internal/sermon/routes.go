package sermon

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group, h *Handler) {

	e.GET("/articles/weekly", h.GetWeeklyArticles)
	e.GET("/articles/monthly", h.GetMonthlyArticles)
	e.GET("/articles/yearly", h.GetYearlyArticles)
	e.GET("/articles/:id", h.GetDetailArticle)
}
