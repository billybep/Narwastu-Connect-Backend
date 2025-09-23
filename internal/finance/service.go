package finance

import (
	"time"
)

type Service interface {
	ListAccounts() ([]Account, error)
	ListTransactions(accountID uint) ([]Transaction, error)
	AddTransaction(tx *Transaction) error
	GetBalance(accountID uint) (int64, error)
	GetWeeklyReport(accountID uint) (*WeeklyReport, error)
	GetWeeklySummary() ([]WeeklySummary, error)
	GetWeeklyTransactions(accountID uint) ([]TransactionDTO, time.Time, time.Time, error)
}

type service struct {
	repo Repository
}

// =========================
// 📌 Weekly Report
// =========================
type WeeklyReport struct {
	AccountID      uint          `json:"account_id"`
	AccountName    string        `json:"account_name"`
	PeriodStart    time.Time     `json:"period_start"`
	PeriodEnd      time.Time     `json:"period_end"`
	OpeningBalance int64         `json:"opening_balance"`
	Transactions   []Transaction `json:"transactions"`
	ClosingBalance int64         `json:"closing_balance"`
}

type TransactionDTO struct {
	ID          uint   `json:"id"`
	AccountID   uint   `json:"account_id"`
	Date        string `json:"date"`
	Type        string `json:"type"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) ListAccounts() ([]Account, error) {
	return s.repo.GetAccounts()
}

func (s *service) ListTransactions(accountID uint) ([]Transaction, error) {
	return s.repo.GetTransactions(accountID)
}

func (s *service) GetBalance(accountID uint) (int64, error) {
	return s.repo.GetBalance(accountID)
}

func (s *service) GetWeeklyReport(accountID uint) (*WeeklyReport, error) {
	now := time.Now()

	// cari range minggu lalu (Minggu–Sabtu)
	weekday := int(now.Weekday()) // 0 = Sunday
	startOfThisWeek := now.AddDate(0, 0, -weekday)
	start := startOfThisWeek.AddDate(0, 0, -7)
	end := start.AddDate(0, 0, 6)

	// ambil nama account
	accounts, err := s.repo.GetAccounts()
	if err != nil {
		return nil, err
	}
	var accountName string
	for _, a := range accounts {
		if a.ID == accountID {
			accountName = a.Name
			break
		}
	}

	// 🔑 ambil saldo real-time dari balances
	openingBalance, err := s.repo.GetBalance(accountID)
	if err != nil {
		return nil, err
	}

	// ambil transaksi minggu ini
	transactions, err := s.repo.FindWeeklyTransactions(accountID, start, end)
	if err != nil {
		return nil, err
	}

	// hitung closing berdasarkan transaksi minggu ini
	closingBalance := openingBalance
	for _, t := range transactions {
		if t.Type == "DEBIT" {
			closingBalance += t.Amount
		} else {
			closingBalance -= t.Amount
		}
	}

	return &WeeklyReport{
		AccountID:      accountID,
		AccountName:    accountName,
		PeriodStart:    start,
		PeriodEnd:      end,
		OpeningBalance: openingBalance,
		Transactions:   transactions,
		ClosingBalance: closingBalance,
	}, nil
}

// response summary per akun
type WeeklySummary struct {
	AccountID      uint   `json:"account_id"`
	AccountName    string `json:"account"`
	Description    string `json:"description"`
	ClosingBalance int64  `json:"closing_balance"`
}

// func (s *service) GetWeeklySummary() ([]WeeklySummary, error) {
// 	now := time.Now()

// 	weekday := int(now.Weekday()) // 0=Sunday
// 	startOfThisWeek := now.AddDate(0, 0, -weekday)
// 	start := startOfThisWeek.AddDate(0, 0, -7)
// 	end := start.AddDate(0, 0, 6)

// 	accounts, err := s.repo.GetAccounts()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var summaries []WeeklySummary
// 	for _, acc := range accounts {
// 		// 🔑 ambil saldo real-time
// 		openingBalance, err := s.repo.GetBalance(acc.ID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// ambil transaksi minggu ini
// 		transactions, err := s.repo.FindWeeklyTransactions(acc.ID, start, end)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// hitung closing balance
// 		closingBalance := openingBalance
// 		for _, t := range transactions {
// 			if t.Type == "DEBIT" {
// 				closingBalance += t.Amount
// 			} else {
// 				closingBalance -= t.Amount
// 			}
// 		}

// 		summaries = append(summaries, WeeklySummary{
// 			AccountID:      acc.ID,
// 			AccountName:    acc.Name,
// 			Description:    acc.Description,
// 			ClosingBalance: closingBalance,
// 		})
// 	}

// 	return summaries, nil
// }

func (s *service) GetWeeklySummary() ([]WeeklySummary, error) {

	accounts, err := s.repo.GetAccounts()
	if err != nil {
		return nil, err
	}

	var summaries []WeeklySummary
	for _, acc := range accounts {
		// 🔑 ambil saldo real-time (langsung dari balances table)
		realBalance, err := s.repo.GetBalance(acc.ID)
		if err != nil {
			return nil, err
		}

		summaries = append(summaries, WeeklySummary{
			AccountID:      acc.ID,
			AccountName:    acc.Name,
			Description:    acc.Description,
			ClosingBalance: realBalance, // langsung pakai saldo nyata
		})
	}

	return summaries, nil
}

// --- hitung range minggu yang baru lewat
func getLastWeekRange() (time.Time, time.Time) {
	now := time.Now()
	weekday := int(now.Weekday())          // 0=Sunday
	start := now.AddDate(0, 0, -weekday-7) // minggu sebelumnya
	end := start.AddDate(0, 0, 6)
	return start, end
}

// func (s *service) GetWeeklyTransactions(accountID uint) ([]Transaction, time.Time, time.Time, error) {
// 	start, end := getLastWeekRange()
// 	txs, err := s.repo.FindWeeklyTransactions(accountID, start, end)
// 	return txs, start, end, err
// }

func (s *service) GetWeeklyTransactions(accountID uint) ([]TransactionDTO, time.Time, time.Time, error) {
	start, end := getLastWeekRange()
	txs, err := s.repo.FindWeeklyTransactions(accountID, start, end)
	if err != nil {
		return nil, start, end, err
	}

	var result []TransactionDTO
	for _, t := range txs {
		result = append(result, TransactionDTO{
			ID:          t.ID,
			AccountID:   t.AccountID,
			Date:        t.Date.Format("2006-01-02"), // hanya YYYY-MM-DD
			Type:        t.Type,
			Amount:      t.Amount,
			Description: t.Description,
		})
	}

	return result, start, end, nil
}

func (s *service) AddTransaction(tx *Transaction) error {
	// simpan transaksi
	err := s.repo.CreateTransaction(tx)
	if err != nil {
		return err
	}

	// hitung perubahan saldo
	var delta int64
	switch tx.Type {
	case "DEBIT":
		delta = tx.Amount
	case "CREDIT":
		delta = -tx.Amount
	}

	// update saldo akun
	return s.repo.(*repository).UpdateBalance(tx.AccountID, delta)
}
