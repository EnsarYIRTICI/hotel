package service

import (
	"errors"
	"go-basic/hotel/dto"
	"go-basic/hotel/entity"
	"go-basic/hotel/repository"

	"gorm.io/gorm"
)

// HotelService, otel işlemlerini yönetir
type HotelService struct {
	DB        *gorm.DB
	HotelRepo *repository.HotelRepository
}

// Yeni HotelService oluşturur
func NewHotelService(db *gorm.DB) *HotelService {
	return &HotelService{
		DB:        db,
		HotelRepo: repository.NewHotelRepository(db),
	}
}

// CreateHotel, yeni bir otel oluşturur
func (s *HotelService) CreateHotel(hotelDTO dto.CreateHotelDTO) (*entity.Hotel, error) {
	// Otel adı boş olamaz
	if hotelDTO.Name == "" {
		return nil, errors.New("Otel adı boş olamaz")
	}

	// Yeni otel nesnesi oluştur
	hotel := &entity.Hotel{
		Name:        hotelDTO.Name,
		Address:     hotelDTO.Address,
		City:        hotelDTO.City,
		Country:     hotelDTO.Country,
		Description: hotelDTO.Description,
	}

	// Transaction başlat
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// Oteli kaydet
		if err := tx.Create(hotel).Error; err != nil {
			return err
		}

		// Eğer otel oluşturulursa, işlem tamamlanır
		return nil
	})

	// Transaction hata kontrolü
	if err != nil {
		return nil, err
	}

	// Başarıyla otel oluşturuldu
	return hotel, nil
}
