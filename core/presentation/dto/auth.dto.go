package dto

// LoginDTO, giriş işlemi için gerekli veriyi taşır
type LoginDTO struct {
	IdentityNumber string `json:"identity_number"`
	Password       string `json:"password"`
}

type RegisterDTO struct {
	IdentityNumber string `json:"identity_number"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}
