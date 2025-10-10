package admin

import (
	"app/internal/finance"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTransaction(tx *finance.Transaction) error
	GetTransactions(accountID uint) ([]finance.Transaction, error)
	UpdateBalance(accountID uint, delta int64) error
	GetAccounts() ([]finance.Account, error)
	GetBalance(accountID uint) (int64, error)
	FindWeeklyTransactions(accountID uint, start, end time.Time) ([]finance.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

// === TRANSAKSI ===
func (r *repository) CreateTransaction(tx *finance.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *repository) GetTransactions(accountID uint) ([]finance.Transaction, error) {
	var txs []finance.Transaction
	err := r.db.Where("account_id = ?", accountID).Order("date DESC").Find(&txs).Error
	return txs, err
}

func (r *repository) UpdateBalance(accountID uint, delta int64) error {
	var bal finance.Balance
	err := r.db.Where("account_id = ?", accountID).First(&bal).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			bal = finance.Balance{
				AccountID:   accountID,
				Balance:     delta,
				LastUpdated: time.Now(),
			}
			return r.db.Create(&bal).Error
		}
		return err
	}

	bal.Balance += delta
	bal.LastUpdated = time.Now()
	return r.db.Save(&bal).Error
}

// === ACCOUNT & BALANCE ===
func (r *repository) GetAccounts() ([]finance.Account, error) {
	var accounts []finance.Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}

func (r *repository) GetBalance(accountID uint) (int64, error) {
	var bal finance.Balance
	err := r.db.Where("account_id = ?", accountID).FirstOrCreate(&bal, finance.Balance{
		AccountID: accountID,
		Balance:   0,
	}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return bal.Balance, nil
}

func (r *repository) FindWeeklyTransactions(accountID uint, start, end time.Time) ([]finance.Transaction, error) {
	var txs []finance.Transaction
	err := r.db.
		Where("account_id = ? AND DATE(date) BETWEEN ? AND ?", accountID, start.Format("2006-01-02"), end.Format("2006-01-02")).
		Order("date ASC").
		Find(&txs).Error
	return txs, err
}
