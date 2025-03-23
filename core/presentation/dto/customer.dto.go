package dto

// CustomerResponseDTO, müşteri bilgilerini döndürmek için kullanılan DTO'dur.
type CustomerResponseDTO struct {
	ID            uint   `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Address       string `json:"address,omitempty"`
	LoyaltyPoints int    `json:"loyalty_points"`
}
