package finance

import (
	"time"

	"gorm.io/gorm"
)

func SeedFinance(db *gorm.DB) error {
	accounts := []Account{
		{Name: "Kas Pembangunan", Description: "Kas Pembangunan Pastori dan Sarana Gereja lainnya"},
		{Name: "Kas Pengadaan Kursi", Description: "Kas untuk pengadaan kursi gereja"},
		{Name: "Kas Gereja", Description: "Kas Gereja utama"},
		{Name: "Kas Diakonia", Description: "Kas Diakonia"},
		{Name: "Kas Sekretariat Gereja", Description: "Kas Sekretariat Gereja"},
		{Name: "Kas Tika Pro", Description: "Kas Tindakan Kasih Prokorus"},
		{Name: "Kas Rukun MKM", Description: "Kas Makan Kasih Maleos-leosan"},
	}

	for _, acc := range accounts {
		db.FirstOrCreate(&acc, Account{Name: acc.Name})
	}

	// ambil account untuk relasi transaksi
	var kasGereja Account
	db.First(&kasGereja, "name = ?", "Kas Gereja")

	// transaksi untuk Kas Gereja minggu 7 Sept 2025
	transactions := []Transaction{
		{
			AccountID:   kasGereja.ID,
			Date:        time.Date(2025, 9, 7, 0, 0, 0, 0, time.Local),
			Type:        "DEBIT",
			Amount:      1101600,
			Description: "Persembahan Ibadah Raya",
		},
		{
			AccountID:   kasGereja.ID,
			Date:        time.Date(2025, 9, 7, 0, 0, 0, 0, time.Local),
			Type:        "CREDIT",
			Amount:      300000,
			Description: "PK Pemain Musik",
		},
		{
			AccountID:   kasGereja.ID,
			Date:        time.Date(2025, 9, 7, 0, 0, 0, 0, time.Local),
			Type:        "CREDIT",
			Amount:      500000,
			Description: "Support Ibadah Pelprap Wilayah",
		},
	}

	for _, tx := range transactions {
		db.Create(&tx)
	}

	// saldo awal per akun (dari gambar)
	initialBalances := map[string]int64{
		"Kas Pembangunan":                           15156000,
		"Kas Pengadaan Kursi":                       23000000,
		"Kas Gereja":                                1788400, // sebelum transaksi 7 Sept
		"Kas Diakonia":                              6050000,
		"Kas Sekretariat Gereja":                    792100,
		"Kas Tika Pro (Tindakan Kasih Prokorus)":    0,
		"Kas Rukun MKM (Makan Kasih Maleos-leosan)": 1370000,
	}

	for name, balance := range initialBalances {
		var acc Account
		if err := db.First(&acc, "name = ?", name).Error; err == nil {
			// simpan saldo awal ke tabel balances
			var bal Balance
			db.FirstOrCreate(&bal, Balance{AccountID: acc.ID}, Balance{
				AccountID:   acc.ID,
				Balance:     balance,
				LastUpdated: time.Now(),
			})
		}
	}

	return nil
}
