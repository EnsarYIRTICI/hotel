package entity

type Hotel struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:200;not null"`
	Address     string  `gorm:"size:300"`
	City        string  `gorm:"size:100"`
	Country     string  `gorm:"size:100"`
	Description string  `gorm:"type:text"`
	Rating      float32 `gorm:"default:0"` // Ortalama puan
	Rooms       []Room  `gorm:"foreignKey:HotelID"`
}
