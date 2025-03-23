package security

import (
	"go-basic/domain/entity"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwtKey []byte // Güvenlik amacıyla gizli bir anahtar

// Token yapısı
type Claims struct {
	IdentityNumber string `json:"identity_number"`
	jwt.StandardClaims
}

// .env dosyasındaki değişkenleri yükle
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// JWT_SECRET_KEY çevresel değişkenini al ve jwtKey'e ata
	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET_KEY not set in .env file")
	}
}

// JWT token oluşturma
func GenerateJWT(customer entity.Customer) string {
	// Token süresi
	expirationTime := time.Now().Add(24 * time.Hour)

	// Claims (Token içeriği)
	claims := &Claims{
		IdentityNumber: customer.IdentityNumber, // IdentityNumber'ı token'a ekledik
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "go-basic-app",
		},
	}

	// Yeni JWT token oluştur
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatalf("Token oluşturulamadı: %v", err)
	}

	return tokenString
}

// JWT doğrulama
func ValidateJWT(tokenString string) (*Claims, error) {
	// Token'ı çözümle
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Algoritma doğrulaması yapılıyor
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	// Token geçerli ise, claims döndürülür
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
