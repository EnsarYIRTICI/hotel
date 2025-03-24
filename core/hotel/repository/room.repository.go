package repository

import (
	"go-basic/base"
	"go-basic/hotel/entity"

	"gorm.io/gorm"
)

// RoomRepository, oda işlemlerini yönetir
type RoomRepository struct {
	base.BaseRepository[entity.Room]
}

// Yeni RoomRepository oluşturur
func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		BaseRepository: *base.NewBaseRepository[entity.Room](db),
	}
}

// Odanın rezerve edilip edilmediğini güncelle
func (r *RoomRepository) UpdateBookingStatus(roomID uint, isBooked bool) error {
	return r.DB.Model(&entity.Room{}).Where("id = ?", roomID).Update("is_booked", isBooked).Error
}
