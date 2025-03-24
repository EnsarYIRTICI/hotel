package entity

import (
	"time"
)

// Customer represents a customer in the system.
type Customer struct {
	ID             uint      `json:"id" gorm:"primaryKey"`          // Benzersiz müşteri kimliği
	FirstName      string    `json:"first_name" gorm:"not null"`    // Müşterinin adı
	LastName       string    `json:"last_name" gorm:"not null"`     // Müşterinin soyadı
	Email          string    `json:"email" gorm:"unique;not null"`  // Müşterinin e-posta adresi
	PhoneNumber    string    `json:"phone_number" gorm:"not null"`  // Müşterinin telefon numarası
	Address        string    `json:"address"`                       // Müşterinin adresi (isteğe bağlı)
	DateOfBirth    time.Time `json:"date_of_birth"`                 // Müşterinin doğum tarihi
	IdentityNumber string    `json:"identity_number" gorm:"unique"` // Müşterinin kimlik numarası
	Password       string    `json:"password" gorm:"not null"`      // Müşterinin şifresi
	LoyaltyPoints  int       `json:"loyalty_points"`                // Müşterinin sadakat puanları (isteğe bağlı)
	CreatedAt      time.Time `json:"created_at"`                    // Kayıt oluşturulma tarihi
	UpdatedAt      time.Time `json:"updated_at"`                    // Kayıt güncellenme tarihi
}
