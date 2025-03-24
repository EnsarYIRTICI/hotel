package repository

import (
	"go-basic/base"
	reservation "go-basic/reservation/entity"

	"gorm.io/gorm"
)

// ReservationRepository, rezervasyon işlemlerini yönetir
type ReservationRepository struct {
	base.BaseRepository[reservation.Reservation]
}

// Yeni ReservationRepository oluşturur
func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{
		BaseRepository: *base.NewBaseRepository[reservation.Reservation](db),
	}
}
