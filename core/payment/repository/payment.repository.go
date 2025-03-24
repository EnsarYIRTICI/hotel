package repository

import (
	"go-basic/base"
	"go-basic/payment/entity"

	"gorm.io/gorm"
)

// PaymentRepository, ödeme işlemlerini yönetir
type PaymentRepository struct {
	base.BaseRepository[entity.Payment]
}

// Yeni PaymentRepository oluşturur
func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		BaseRepository: *base.NewBaseRepository[entity.Payment](db),
	}
}
