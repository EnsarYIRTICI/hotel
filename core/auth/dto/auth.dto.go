package dto

type LoginDTO struct {
	IdentityNumber string `json:"identity_number"`
	Password       string `json:"password"`
}

type RegisterDTO struct {
	IdentityNumber string `json:"identity_number"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

type AuthResponseDTO struct {
	ID            uint   `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Address       string `json:"address,omitempty"`
	LoyaltyPoints int    `json:"loyalty_points"`
}
