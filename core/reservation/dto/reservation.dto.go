package dto

import "time"

// CreateReservationDTO, rezervasyon oluşturma isteği için gerekli verileri taşır.
type CreateReservationDTO struct {
	CustomerID uint      `json:"customer_id" binding:"required"`
	RoomID     uint      `json:"room_id" binding:"required"`
	CheckIn    time.Time `json:"check_in" binding:"required"`
	CheckOut   time.Time `json:"check_out" binding:"required"`
}
