package dto

// CreateRoomDTO, yeni bir oda oluşturmak için gerekli verileri taşır
type CreateRoomDTO struct {
	HotelID  uint    `json:"hotel_id" binding:"required"`
	Type     string  `json:"type" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Capacity int     `json:"capacity" binding:"required"`
}

// UpdateRoomDTO, mevcut bir odayı güncellemek için gerekli verileri taşır
type UpdateRoomDTO struct {
	ID       uint    `json:"id" binding:"required"`
	Type     string  `json:"type" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Capacity int     `json:"capacity" binding:"required"`
}
