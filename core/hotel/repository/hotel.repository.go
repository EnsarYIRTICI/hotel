package repository

import (
	"go-basic/base"
	"go-basic/hotel/entity"

	"gorm.io/gorm"
)

// HotelRepository, otel işlemlerini yönetir
type HotelRepository struct {
	base.BaseRepository[entity.Hotel]
}

// Yeni HotelRepository oluşturur
func NewHotelRepository(db *gorm.DB) *HotelRepository {
	return &HotelRepository{
		BaseRepository: *base.NewBaseRepository[entity.Hotel](db),
	}
}
