package finance

import (
	"time"
)

type Account struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	AccountID   uint      `json:"account_id"`
	Account     Account   `gorm:"foreignKey:AccountID" json:"account"`
	Date        time.Time `json:"date"`
	Type        string    `gorm:"size:10;check:type IN ('DEBIT','CREDIT')" json:"type"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Balance struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	AccountID   uint      `json:"account_id"`
	Account     Account   `gorm:"foreignKey:AccountID" json:"account"`
	Balance     int64     `json:"balance"`
	LastUpdated time.Time `json:"last_updated"`
}
