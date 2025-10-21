package finance

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) GetAccounts(c echo.Context) error {
	accounts, err := h.svc.ListAccounts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, accounts)
}

func (h *Handler) GetTransactions(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transactions, err := h.svc.ListTransactions(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactions)
}

func (h *Handler) CreateTransaction(c echo.Context) error {
	var tx Transaction
	if err := c.Bind(&tx); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.svc.AddTransaction(&tx); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, tx)
}

func (h *Handler) GetBalance(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	balance, err := h.svc.GetBalance(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"account_id": id,
		"balance":    balance,
	})
}

// ======================
// 📌 Weekly Report
// ======================
func (h *Handler) GetWeeklyReport(c echo.Context) error {
	accountID, _ := strconv.Atoi(c.QueryParam("account_id"))

	report, err := h.svc.GetWeeklyReport(uint(accountID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, report)
}

func (h *Handler) GetWeeklySummary(c echo.Context) error {
	data, err := h.svc.GetWeeklySummary()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) GetWeeklyTransactions(c echo.Context) error {
	idStr := c.Param("id")
	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid account id"})
	}

	txs, start, end, err := h.svc.GetWeeklyTransactions(uint(accountID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"range":        start.Format("2006-01-02") + " s/d " + end.Format("2006-01-02"),
		"transactions": txs,
	})
}
