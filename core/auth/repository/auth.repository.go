package repository

import (
	"errors"
	"go-basic/customer/entity"
	"go-basic/infrastructure/db"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(customer *entity.Customer) error
	GetByIdentityNumber(IdentityNumber string) (*entity.Customer, error)
}

type AuthRepositoryST struct {
	db *gorm.DB
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryST{db: db.DB}
}

func (r *AuthRepositoryST) Create(customer *entity.Customer) error {
	return r.db.Create(customer).Error
}

func (r *AuthRepositoryST) GetByIdentityNumber(IdentityNumber string) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Where("identity_number = ?", IdentityNumber).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &customer, err
}
