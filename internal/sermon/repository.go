package sermon

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetMonthlyArticles() ([]Sermon, error)
	GetByID(id uint) (*Sermon, error)
	FindLatestArticles(limit int) ([]Sermon, error)
	FindArticlesByDateRange(start, end time.Time) ([]Sermon, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetMonthlyArticles() ([]Sermon, error) {
	var sermons []Sermon
	// Ambil 4 artikel terbaru bulan ini
	err := r.db.Where("date >= date_trunc('month', CURRENT_DATE)").
		Order("date desc").
		Limit(4).
		Find(&sermons).Error
	return sermons, err
}

func (r *repository) GetByID(id uint) (*Sermon, error) {
	var sermon Sermon
	err := r.db.First(&sermon, id).Error
	if err != nil {
		return nil, err
	}
	return &sermon, nil
}

func (r *repository) FindLatestArticles(limit int) ([]Sermon, error) {
	var articles []Sermon
	err := r.db.Order("date DESC").Limit(limit).Find(&articles).Error
	return articles, err
}

func (r *repository) FindArticlesByDateRange(start, end time.Time) ([]Sermon, error) {
	var articles []Sermon
	err := r.db.Where("date BETWEEN ? AND ?", start, end).
		Order("date DESC").
		Find(&articles).Error
	return articles, err
}
