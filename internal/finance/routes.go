package finance

import "github.com/labstack/echo/v4"

func RegisterRoutes(g *echo.Group, h *Handler) {
	g.GET("/accounts", h.GetAccounts)
	g.GET("/accounts/:id/transactions", h.GetTransactions)
	g.POST("/transactions", h.CreateTransaction)
	g.GET("/accounts/:id/balance", h.GetBalance)

	// 📌 laporan mingguan
	g.GET("/reports/finance/weekly", h.GetWeeklyReport)

	// list saldo mingguan
	g.GET("/reports/finance/weekly/summary", h.GetWeeklySummary)
	// detail transaksi mingguan per akun
	g.GET("/reports/finance/weekly/accounts/:id", h.GetWeeklyTransactions)
}
