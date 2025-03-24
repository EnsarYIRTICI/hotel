package entity

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	ReservationID uint      `gorm:"not null"`
	Amount        float64   `gorm:"not null"`
	PaymentDate   time.Time `gorm:"autoCreateTime"`
	Status        string    `gorm:"size:50;default:'pending'"` // pending, completed, failed
}
