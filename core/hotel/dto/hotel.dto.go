package dto

// CreateHotelDTO, otel oluşturma işlemi için gerekli verileri taşır
type CreateHotelDTO struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Description string `json:"description"`
}
