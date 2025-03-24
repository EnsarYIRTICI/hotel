package repository

import (
	"go-basic/base"
	review "go-basic/review/entity" // HL

	"gorm.io/gorm"
)

// ReviewRepository, yorum işlemlerini yönetir
type ReviewRepository struct {
	base.BaseRepository[review.Review]
}

// Yeni ReviewRepository oluşturur
func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{
		BaseRepository: *base.NewBaseRepository[review.Review](db),
	}
}
