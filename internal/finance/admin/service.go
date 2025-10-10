package admin

import (
	"app/internal/finance"
	"errors"
	"time"
)

type Service interface {
	CreateIncome(accountID uint, amount int64, description string) error
	CreateExpense(accountID uint, amount int64, description string) error
	GetTransactionsByAccount(accountID uint) ([]finance.Transaction, error)
	GetWeeklySummary() ([]finance.WeeklySummary, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// === CREATE INCOME ===
func (s *service) CreateIncome(accountID uint, amount int64, description string) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	tx := &finance.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Type:        "DEBIT",
		Description: description,
		Date:        time.Now(),
	}

	return s.addTransaction(tx)
}

// === CREATE EXPENSE ===
func (s *service) CreateExpense(accountID uint, amount int64, description string) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	tx := &finance.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Type:        "CREDIT",
		Description: description,
		Date:        time.Now(),
	}

	return s.addTransaction(tx)
}

// === PRIVATE: add transaction + update balance ===
func (s *service) addTransaction(tx *finance.Transaction) error {
	if err := s.repo.CreateTransaction(tx); err != nil {
		return err
	}

	var delta int64
	switch tx.Type {
	case "DEBIT":
		delta = tx.Amount
	case "CREDIT":
		delta = -tx.Amount
	}

	return s.repo.UpdateBalance(tx.AccountID, delta)
}

// === GET TRANSAKSI PER ACCOUNT ===
func (s *service) GetTransactionsByAccount(accountID uint) ([]finance.Transaction, error) {
	return s.repo.GetTransactions(accountID)
}

// === GET WEEKLY SUMMARY ===
func (s *service) GetWeeklySummary() ([]finance.WeeklySummary, error) {
	accounts, err := s.repo.GetAccounts()
	if err != nil {
		return nil, err
	}

	var summaries []finance.WeeklySummary
	for _, acc := range accounts {
		balance, err := s.repo.GetBalance(acc.ID)
		if err != nil {
			return nil, err
		}

		summaries = append(summaries, finance.WeeklySummary{
			AccountID:      acc.ID,
			AccountName:    acc.Name,
			Description:    acc.Description,
			ClosingBalance: balance,
		})
	}

	return summaries, nil
}
