package entity

import (
	customer "go-basic/customer/entity" // HL
	hotel "go-basic/hotel/entity"       // HL
	payment "go-basic/payment/entity"   // HL
	review "go-basic/review/entity"     // HL
	"time"
)

type Reservation struct {
	ID         uint              `gorm:"primaryKey"`
	CustomerID uint              `gorm:"not null"`
	Customer   customer.Customer `gorm:"foreignKey:CustomerID"`

	RoomID uint       `gorm:"not null"`
	Room   hotel.Room `gorm:"foreignKey:RoomID"`

	CheckIn  time.Time `gorm:"not null"`
	CheckOut time.Time `gorm:"not null"`

	Status string `gorm:"size:50;default:'pending'"` // pending, confirmed, canceled

	Payment payment.Payment `gorm:"foreignKey:ReservationID"`
	Review  review.Review   `gorm:"foreignKey:ReservationID"`
}
