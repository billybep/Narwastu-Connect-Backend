package admin

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

type transactionInput struct {
	AccountID   uint   `json:"account_id"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

func (h *Handler) CreateIncome(c echo.Context) error {
	var input transactionInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	if err := h.service.CreateIncome(input.AccountID, input.Amount, input.Description); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "income transaction created"})
}

func (h *Handler) CreateExpense(c echo.Context) error {
	var input transactionInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	if err := h.service.CreateExpense(input.AccountID, input.Amount, input.Description); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "expense transaction created"})
}

func (h *Handler) GetTransactionsByAccount(c echo.Context) error {
	accountIDParam := c.Param("accountID")
	accountID, err := strconv.Atoi(accountIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid account id"})
	}

	txs, err := h.service.GetTransactionsByAccount(uint(accountID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, txs)
}

func (h *Handler) GetWeeklySummary(c echo.Context) error {
	data, err := h.service.GetWeeklySummary()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}
