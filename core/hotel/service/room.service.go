package service

import (
	"errors"
	"go-basic/hotel/dto"
	"go-basic/hotel/entity"
	"go-basic/hotel/repository"

	"gorm.io/gorm"
)

// RoomService, oda işlemlerini yönetir
type RoomService struct {
	DB       *gorm.DB
	RoomRepo *repository.RoomRepository
}

// Yeni RoomService oluşturur
func NewRoomService(db *gorm.DB) *RoomService {
	return &RoomService{
		DB:       db,
		RoomRepo: repository.NewRoomRepository(db),
	}
}

// CreateRoom, yeni bir oda oluşturur
func (s *RoomService) CreateRoom(roomDTO dto.CreateRoomDTO) (*entity.Room, error) {
	// Oda fiyatı 0 olamaz
	if roomDTO.Price <= 0 {
		return nil, errors.New("Oda fiyatı geçerli olmalıdır")
	}

	// Oda tipi boş olamaz
	if roomDTO.Type == "" {
		return nil, errors.New("Oda tipi boş olamaz")
	}

	// Oda kapasitesi 0 olamaz
	if roomDTO.Capacity <= 0 {
		return nil, errors.New("Oda kapasitesi geçerli olmalıdır")
	}

	// Yeni oda nesnesi oluştur
	room := &entity.Room{
		HotelID:  roomDTO.HotelID,
		Type:     roomDTO.Type,
		Price:    roomDTO.Price,
		Capacity: roomDTO.Capacity,
		IsBooked: false, // Yeni oda başlangıçta müsait
	}

	// Transaction başlat
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// Odayı kaydet
		if err := tx.Create(room).Error; err != nil {
			return err
		}

		return nil
	})

	// Transaction hata kontrolü
	if err != nil {
		return nil, err
	}

	// Başarıyla oda oluşturuldu
	return room, nil
}

// UpdateRoom, mevcut bir odayı günceller
func (s *RoomService) UpdateRoom(roomDTO dto.UpdateRoomDTO) (*entity.Room, error) {
	// Oda ID'si gerekli
	if roomDTO.ID == 0 {
		return nil, errors.New("Geçerli bir oda ID'si gerekli")
	}

	var room entity.Room
	if err := s.DB.First(&room, roomDTO.ID).Error; err != nil {
		return nil, errors.New("Oda bulunamadı")
	}

	// Odayı güncelle
	room.Type = roomDTO.Type
	room.Price = roomDTO.Price
	room.Capacity = roomDTO.Capacity

	// Transaction başlat
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		// Odayı kaydet
		if err := tx.Save(&room).Error; err != nil {
			return err
		}
		return nil
	})

	// Transaction hata kontrolü
	if err != nil {
		return nil, err
	}

	// Başarıyla oda güncellendi
	return &room, nil
}

// DeleteRoom, odayı siler
func (s *RoomService) DeleteRoom(roomID uint) error {
	if roomID == 0 {
		return errors.New("Geçerli bir oda ID'si gerekli")
	}

	var room entity.Room
	if err := s.DB.First(&room, roomID).Error; err != nil {
		return errors.New("Oda bulunamadı")
	}

	// Odayı sil
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&room).Error; err != nil {
			return err
		}
		return nil
	})

	// Transaction hata kontrolü
	if err != nil {
		return err
	}

	return nil
}
