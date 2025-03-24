package entity

import "time"

type Review struct {
	ID         uint      `gorm:"primaryKey"`
	CustomerID uint      `gorm:"not null"`
	HotelID    uint      `gorm:"not null"`
	Rating     int       `gorm:"not null"` // 1-5 arasında puan
	Comment    string    `gorm:"type:text"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
