package finance

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetAccounts() ([]Account, error)
	GetTransactions(accountID uint) ([]Transaction, error)
	CreateTransaction(tx *Transaction) error
	GetBalance(accountID uint) (int64, error)
	FindWeeklyTransactions(accountID uint, start, end time.Time) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAccounts() ([]Account, error) {
	var accounts []Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}

func (r *repository) GetTransactions(accountID uint) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("account_id = ?", accountID).Order("date DESC").Find(&transactions).Error
	return transactions, err
}

func (r *repository) CreateTransaction(tx *Transaction) error {
	return r.db.Create(tx).Error
}

func (r *repository) GetBalance(accountID uint) (int64, error) {
	var bal Balance
	err := r.db.Where("account_id = ?", accountID).FirstOrCreate(&bal, Balance{
		AccountID: accountID,
		Balance:   0,
	}).Error

	if err != nil {
		// kalau belum ada record balance, default 0
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return bal.Balance, nil
}

// --- Weekly transactions per account
func (r *repository) FindWeeklyTransactions(accountID uint, start, end time.Time) ([]Transaction, error) {
	var txs []Transaction
	err := r.db.
		Where("account_id = ? AND DATE(date) BETWEEN ? AND ?", accountID,
			start.Format("2006-01-02"), end.Format("2006-01-02")).
		Order("date ASC").
		Find(&txs).Error
	return txs, err
}

// Simpan / update saldo saat ada transaksi
func (r *repository) UpdateBalance(accountID uint, delta int64) error {
	var bal Balance
	err := r.db.Where("account_id = ?", accountID).First(&bal).Error
	if err != nil {
		// kalau belum ada, insert baru
		if err == gorm.ErrRecordNotFound {
			bal = Balance{
				AccountID:   accountID,
				Balance:     delta,
				LastUpdated: time.Now(),
			}
			return r.db.Create(&bal).Error
		}
		return err
	}

	// update saldo existing
	bal.Balance += delta
	bal.LastUpdated = time.Now()
	return r.db.Save(&bal).Error
}
