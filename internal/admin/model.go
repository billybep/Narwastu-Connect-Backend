package admin

import "time"

type Role string

const (
	RoleSystemAdmin Role = "system_administrator"
	RoleAdmin       Role = "administrator"
	RoleMember      Role = "member"
	RoleGuest       Role = "guest"
)

type Admin struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email" gorm:"uniqueIndex;not null"`
	Password    string    `json:"-"` // hash
	Role        Role      `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
