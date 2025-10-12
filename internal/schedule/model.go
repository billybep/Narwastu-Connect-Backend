package schedule

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ServiceSchedule struct {
	ID     uint      `json:"id" gorm:"primaryKey"`
	Date   time.Time `json:"date"`
	Leader string    `json:"leader"`

	Singers       datatypes.JSON `json:"singers"`
	Tambourines   datatypes.JSON `json:"tambourines"`
	Banners       datatypes.JSON `json:"banners"`
	Musicians     datatypes.JSON `json:"musicians"`
	Multimedia    datatypes.JSON `json:"multimedia"`
	Collectors    datatypes.JSON `json:"collectors"`
	WorshipGroups datatypes.JSON `json:"worshipGroups"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
