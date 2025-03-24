package entity

type Room struct {
	ID       uint    `gorm:"primaryKey"`
	HotelID  uint    `gorm:"not null"`
	Type     string  `gorm:"size:50;not null"` // Single, Double, Suite vs.
	Price    float64 `gorm:"not null"`
	Capacity int     `gorm:"not null"`      // Kaç kişi kalabilir?
	IsBooked bool    `gorm:"default:false"` // Oda müsait mi?
}
