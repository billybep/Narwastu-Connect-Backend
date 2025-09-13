package member

import (
	"time"

	"gorm.io/gorm"
)

type Site struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:50;unique;not null" json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Member struct {
	ID uint `gorm:"primaryKey" json:"id"`
	// NIK           string         `gorm:"size:20;unique" json:"nik"`
	NIK           *string        `gorm:"uniqueIndex;default:null"`
	FullName      string         `gorm:"size:100;not null" json:"fullName"`
	FamilyName    string         `gorm:"size:100"`
	PlaceOfBirth  string         `gorm:"size:50" json:"placeOfBirth"`
	DateOfBirth   time.Time      `gorm:"not null" json:"dateOfBirth"`
	Gender        string         `gorm:"size:10" json:"gender"`
	Address       string         `json:"address"`
	Phone         string         `gorm:"size:20" json:"phone"`
	Email         string         `gorm:"size:100" json:"email"`
	BaptismDate   *time.Time     `json:"baptismDate"`
	SiteID        uint           `json:"siteId"`
	Site          Site           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"site"`
	MaritalStatus string         `gorm:"size:20" json:"maritalStatus"`
	PhotoURL      string         `json:"photoUrl"`
	IsActive      bool           `gorm:"default:true" json:"isActive"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type FamilyRelationship struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	MemberID         uint           `json:"memberId"`
	Member           Member         `gorm:"constraint:OnDelete:CASCADE;" json:"member"`
	FamilyMemberID   uint           `json:"familyMemberId"`
	FamilyMember     Member         `gorm:"constraint:OnDelete:CASCADE;" json:"familyMember"`
	RelationshipType string         `gorm:"size:50;not null" json:"relationshipType"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
