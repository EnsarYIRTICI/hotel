package base

import "gorm.io/gorm"

// BaseRepository, tüm entity'ler için ortak repository fonksiyonlarını içerir.
type BaseRepository[T any] struct {
	DB *gorm.DB
}

// Yeni repository oluşturur
func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db}
}

// Yeni kayıt oluşturur
func (r *BaseRepository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

// ID'ye göre kayıt bulur
func (r *BaseRepository[T]) FindByID(id uint, entity *T) error {
	return r.DB.First(entity, id).Error
}

// Tüm kayıtları listeler
func (r *BaseRepository[T]) FindAll(entities *[]T) error {
	return r.DB.Find(entities).Error
}

// Günceller
func (r *BaseRepository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

// Siler
func (r *BaseRepository[T]) Delete(entity *T) error {
	return r.DB.Delete(entity).Error
}
