package service

import (
	"errors"

	"go-basic/customer/entity"

	"go-basic/auth/dto"
	"go-basic/auth/repository"

	"go-basic/security"
	"go-basic/util"

	"github.com/jinzhu/copier"

	"golang.org/x/crypto/bcrypt"
)

// NewauthService, yeni bir authService oluşturur
func NewAuthService() AuthService {

	return &AuthServiceST{
		repo:   repository.NewAuthRepository(),
		logger: *util.NewLogger("AUTH-SERVICE"),
	}
}

// authService, kimlik doğrulama işlemlerini içerir

type AuthServiceST struct {
	repo   repository.AuthRepository
	logger util.Logger
}

type AuthService interface {
	Login(loginData dto.LoginDTO) (*string, *dto.AuthResponseDTO, error)
	Register(registerDto dto.RegisterDTO) (*string, *dto.AuthResponseDTO, error)
}

// Login kullanıcı girişi yapar ve JWT token döner
func (service *AuthServiceST) Login(loginData dto.LoginDTO) (*string, *dto.AuthResponseDTO, error) {

	service.logger.Log("Deneme")

	// Kullanıcıyı kimlik numarasına göre sorgula
	existingCustomer, err := service.repo.GetByIdentityNumber(loginData.IdentityNumber)

	service.logger.Log("------>", existingCustomer, "<------")

	if err != nil {
		return nil, nil, errors.New("hata oluştu: kullanıcı sorgulanırken bir sorun meydana geldi")
	}

	service.logger.Log("Deneme 2")

	// Check if existingCustomer is nil
	if existingCustomer == nil {
		return nil, nil, errors.New("bu kimlik numarası ile kayıtlı kullanıcı yok")
	}

	// Şifreyi kontrol et
	err = bcrypt.CompareHashAndPassword([]byte(existingCustomer.Password), []byte(loginData.Password))
	if err != nil {
		return nil, nil, errors.New("geçersiz kimlik numarası veya şifre")
	}

	// Entity Dto Dönüşümü
	var customerDTO dto.AuthResponseDTO
	err = copier.Copy(&customerDTO, &existingCustomer)
	if err != nil {
		return nil, nil, errors.New("DTO kopyalama hatası")
	}

	// JWT token oluştur
	token := security.GenerateJWT(*existingCustomer)

	return &token, &customerDTO, nil
}

// Register kullanıcı kaydını gerçekleştirir
func (service *AuthServiceST) Register(registerDto dto.RegisterDTO) (*string, *dto.AuthResponseDTO, error) {
	// Aynı kimlik numarası ile kayıtlı kullanıcı var mı?
	existingCustomer, _ := service.repo.GetByIdentityNumber(registerDto.IdentityNumber)

	service.logger.Log("------>", existingCustomer, "<------")

	if existingCustomer != nil {
		return nil, nil, errors.New("bu kimlik numarası ile kayıtlı kullanıcı zaten var")
	}

	// Şifreyi hash'le
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, errors.New("şifre hash'lenemedi")
	}
	registerDto.Password = string(hashedPassword)

	// Veritabanına kaydetmek için entity oluştur
	var customerEntity entity.Customer
	err = copier.Copy(&customerEntity, &registerDto)
	if err != nil {
		return nil, nil, errors.New("copy hatası")
	}

	// Veritabanına kaydet
	err = service.repo.Create(&customerEntity)
	if err != nil {
		return nil, nil, errors.New("kullanici kaydedilemedi")
	}

	// Aynı kimlik numarası ile kayıtlı kullanıcı var mı?
	existingCustomer, _ = service.repo.GetByIdentityNumber(registerDto.IdentityNumber)
	if existingCustomer == nil {
		return nil, nil, errors.New("kullanıcı bulunamadı")
	}

	// Entity Dto Dönüşümü
	var customerDTO dto.AuthResponseDTO
	err = copier.Copy(&customerDTO, &existingCustomer)
	if err != nil {
		return nil, nil, errors.New("copy hatasi")
	}

	// JWT token oluştur
	token := security.GenerateJWT(*existingCustomer)

	return &token, &customerDTO, err

}
