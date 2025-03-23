package repository

import (
	"errors"
	"go-basic/domain/entity"
	"go-basic/infrastructure/db"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *entity.Customer) error
	GetByID(id uint) (*entity.Customer, error)
	GetByIdentityNumber(IdentityNumber string) (*entity.Customer, error)
	GetAll() ([]entity.Customer, error)
	Update(customer *entity.Customer) error
	Delete(id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{db: db.DB}
}

func (r *customerRepository) Create(customer *entity.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) GetByID(id uint) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.First(&customer, id).Error
	return &customer, err
}

func (r *customerRepository) GetByIdentityNumber(IdentityNumber string) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Where("identity_number = ?", IdentityNumber).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &customer, err
}

func (r *customerRepository) GetAll() ([]entity.Customer, error) {
	var customers []entity.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *customerRepository) Update(customer *entity.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Customer{}, id).Error
}
