package member

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func SeedMembers(db *gorm.DB) error {
	// Ambil site (rayon) untuk mapping
	var karmel, zaitun Site
	if err := db.Where("name = ?", "Karmel").First(&karmel).Error; err != nil {
		return fmt.Errorf("site Karmel not found: %w", err)
	}
	if err := db.Where("name = ?", "Zaitun").First(&zaitun).Error; err != nil {
		return fmt.Errorf("site Zaitun not found: %w", err)
	}

	// Dummy members dengan FamilyName
	members := []Member{
		{
			NIK:         "DUMMY-001",
			FullName:    "Saranesa Sinaga",
			FamilyName:  "Kel. Sinaga-Rawung",
			DateOfBirth: time.Date(1995, 9, 8, 0, 0, 0, 0, time.UTC),
			SiteID:      &karmel.ID,
			PhotoURL:    "https://randomuser.me/api/portraits/women/45.jpg",
		},
		{
			NIK:         "DUMMY-002",
			FullName:    "Exel Lumintang",
			FamilyName:  "Kel. Lumintang-Sonith",
			DateOfBirth: time.Date(1992, 9, 9, 0, 0, 0, 0, time.UTC),
			SiteID:      &zaitun.ID,
			PhotoURL:    "https://randomuser.me/api/portraits/men/44.jpg",
		},
		{
			NIK:         "DUMMY-003",
			FullName:    "Andreas Rawung",
			FamilyName:  "Kel. Rawung-Lumintang",
			DateOfBirth: time.Date(1987, 9, 14, 0, 0, 0, 0, time.UTC),
			SiteID:      &zaitun.ID,
			PhotoURL:    "https://randomuser.me/api/portraits/men/46.jpg",
		},
	}

	// Insert jika belum ada
	for _, m := range members {
		var exists int64
		db.Model(&Member{}).Where("nik = ?", m.NIK).Count(&exists)
		if exists == 0 {
			if err := db.Create(&m).Error; err != nil {
				return fmt.Errorf("failed seed member %s: %w", m.FullName, err)
			}
		}
	}

	return nil
}
