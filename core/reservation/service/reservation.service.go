package service

import (
	"errors"
	hotel "go-basic/hotel/entity"
	pay "go-basic/payment/entity"
	"go-basic/reservation/dto"
	"go-basic/reservation/entity"

	"gorm.io/gorm"
)

// ReservationService, rezervasyon işlemlerini yönetir
type ReservationService struct {
	DB *gorm.DB
}

// Yeni ReservationService oluştur
func NewReservationService(db *gorm.DB) *ReservationService {
	return &ReservationService{
		DB: db,
	}
}

// DTO ile rezervasyon oluşturma (Transaction ile)
func (s *ReservationService) CreateReservation(data dto.CreateReservationDTO, amount float64) (*entity.Reservation, error) {
	var reservation *entity.Reservation

	// Transaction başlat
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// Oda durumunu kontrol et
		var room hotel.Room
		if err := tx.First(&room, data.RoomID).Error; err != nil {
			return errors.New("Oda bulunamadı")
		}
		if room.IsBooked {
			return errors.New("Oda zaten rezerve edilmiş")
		}

		// Rezervasyon oluştur
		reservation = &entity.Reservation{
			CustomerID: data.CustomerID,
			RoomID:     data.RoomID,
			CheckIn:    data.CheckIn,
			CheckOut:   data.CheckOut,
			Status:     "pending",
		}
		if err := tx.Create(reservation).Error; err != nil {
			return err
		}

		// Ödeme kaydını ekle
		payment := &pay.Payment{
			ReservationID: reservation.ID,
			Amount:        amount,
			Status:        "pending",
		}
		if err := tx.Create(payment).Error; err != nil {
			return err
		}

		// Odayı "rezerve" olarak işaretle
		if err := tx.Model(&hotel.Room{}).Where("id = ?", data.RoomID).Update("is_booked", true).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return reservation, nil
}
