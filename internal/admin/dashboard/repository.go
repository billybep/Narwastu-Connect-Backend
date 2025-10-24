package admin

import (
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

type DashboardStats struct {
	TotalMembers    int64   `json:"jemaat"`
	ActiveMembers   int64   `json:"jemaatAktif"`
	TotalFinance    float64 `json:"totalKeuangan"`
	IbadahMingguIni int64   `json:"ibadahMingguIni"`
}

// Get dashboard summary
func (r *Repository) GetStats() (DashboardStats, error) {
	var stats DashboardStats

	// 🔹 Total Member
	r.db.Table("members").Count(&stats.TotalMembers)

	// 🔹 Member aktif
	r.db.Table("members").Where("is_active = ?", true).Count(&stats.ActiveMembers)

	// 🔹 Total saldo keuangan (sum balance)
	r.db.Table("balances").Select("COALESCE(SUM(balance), 0)").Scan(&stats.TotalFinance)

	// 🔹 Ibadah minggu ini
	now := time.Now()
	startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)
	r.db.Table("events").Where("date_time BETWEEN ? AND ?", startOfWeek, endOfWeek).Count(&stats.IbadahMingguIni)

	return stats, nil
}

// Get trend keuangan per bulan (6 bulan terakhir)
func (r *Repository) GetFinanceTrend() ([]map[string]interface{}, error) {
	var trend []map[string]interface{}
	err := r.db.Raw(`
		SELECT 
			TO_CHAR(date, 'Mon') AS month,
			SUM(CASE WHEN type = 'CREDIT' THEN amount ELSE 0 END) AS income,
			SUM(CASE WHEN type = 'DEBIT' THEN amount ELSE 0 END) AS expense
		FROM transactions
		WHERE date >= NOW() - INTERVAL '6 MONTH'
		GROUP BY month
		ORDER BY MIN(date)
	`).Scan(&trend).Error
	return trend, err
}

// ✅ Ambil jadwal ibadah hanya untuk HARI INI (berdasarkan zona waktu WITA)
func (r *Repository) GetSchedule() ([]map[string]interface{}, error) {
	var schedules []map[string]interface{}
	err := r.db.Raw(`
		SELECT 
			TO_CHAR((date_time AT TIME ZONE 'Asia/Makassar'), 'DD Mon YYYY') AS date,
			TO_CHAR((date_time AT TIME ZONE 'Asia/Makassar'), 'HH24:MI') AS time,
			title AS event,
			location,
			description
		FROM events
		WHERE DATE(date_time AT TIME ZONE 'Asia/Makassar') = DATE(NOW() AT TIME ZONE 'Asia/Makassar')
		ORDER BY date_time ASC
	`).Scan(&schedules).Error
	return schedules, err
}

// Get admin list
func (r *Repository) GetAdmins() ([]map[string]interface{}, error) {
	var admins []map[string]interface{}
	err := r.db.Raw(`
		SELECT 
			name,
			role,
			TO_CHAR(updated_at, 'DD Mon YYYY') AS lastLogin
		FROM admins
		ORDER BY updated_at DESC
		LIMIT 10
	`).Scan(&admins).Error
	return admins, err
}
