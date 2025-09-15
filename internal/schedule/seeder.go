package schedule

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func SeedSchedules(db *gorm.DB) error {
	schedules := []ServiceSchedule{
		{
			Date:   time.Date(2025, 9, 21, 0, 0, 0, 0, time.UTC),
			Leader: "Eiva Sumeke D.",
			Singers: datatypes.JSON([]byte(`[
				"Pdm. Bobby Rumagit",
				"Youke Sambuaga",
				"Rinny Latun Dien",
				"Yully Sakul M.",
				"Nova Montung S."
			]`)),

			Tambourines: datatypes.JSON([]byte(`[
				"Kherren Kandouw",
				"Junia Sinaga",
				"Jessi Sumenge"
			]`)),
			Banners: datatypes.JSON([]byte(`[
				"Leandrito Kandouw",
				"Erlangga Sumenge"
			]`)),
			Musicians: datatypes.JSON([]byte(`[
				"Billy Pesoth",
				"Erwin Tumilantouw",
				"Yesaya Kurumbatu"
			]`)),
			Multimedia: datatypes.JSON([]byte(`[
				"Leandrito Kandouw",
				"Imanuel Saraun",
				"Canafaro Sambuaga",
				"Jeremy Legoh"
			]`)),
			Collectors: datatypes.JSON([]byte(`[
				"Tineke Pakasi K.",
				"Aneke Mandolang P.",
				"Selvi Lumintang S.",
				"Yuneke Moningkey"
			]`)),
			WorshipGroups: datatypes.JSON([]byte(`[
				"Rayon Zaitun",
				"KJ PELWAP",
				"Pelayanan Usia Anugerah"
			]`)),
		},
	}

	for _, s := range schedules {
		var exists int64
		db.Model(&ServiceSchedule{}).
			Where("date = ?", s.Date).
			Count(&exists)

		if exists == 0 {
			if err := db.Create(&s).Error; err != nil {
				return fmt.Errorf("failed seed schedule %s: %w", s.Date.Format("2006-01-02"), err)
			}
		}
	}

	return nil
}
