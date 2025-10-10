package admin

import (
	adminmiddleware "app/internal/admin" // ✅ sesuaikan dengan module kamu

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, h *Handler) {
	// === Report routes (accessible by all authenticated users) ===
	reportGroup := g.Group(
		"/report/finance",
		adminmiddleware.AdminJWTMiddleware, // hanya butuh JWT, tanpa role restriction
	)
	reportGroup.GET("/transactions/:accountID", h.GetTransactionsByAccount)
	reportGroup.GET("/summary", h.GetWeeklySummary)

	// Apply middleware khusus admin
	adminGroup := g.Group(
		"/finance",
		adminmiddleware.AdminJWTMiddleware,
		adminmiddleware.RequireAdminRole("system_administrator", "administrator"))

	adminGroup.POST("/income", h.CreateIncome)
	adminGroup.POST("/expense", h.CreateExpense)
}
