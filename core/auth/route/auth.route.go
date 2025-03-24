package route

import (
	"go-basic/auth/dto"
	"go-basic/auth/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Tüm Auth route'larını kaydeder
func RegisterRoutes(router *gin.Engine) {

	st := AuthRouteST{
		_service: service.NewAuthService(),
	}

	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("register", st.Register)
		authGroup.POST("login", st.Login)
	}
}

// Tüm Auth Route Struct
type AuthRouteST struct {
	_service service.AuthService
}

// Login, kullanıcı girişini başlatır
func (st *AuthRouteST) Login(c *gin.Context) {
	var loginData dto.LoginDTO

	// JSON verisini al
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri formatı"})
		return
	}

	// AuthService ile işlemi gerçekleştir
	token, customerRes, err := st._service.Login(loginData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Başarılı giriş, token ve müşteri verisi döndürülür
	c.JSON(http.StatusOK, gin.H{
		"customer": customerRes,
		"token":    token,
	})
}

// Register, kullanıcı kaydını başlatır
func (st *AuthRouteST) Register(c *gin.Context) {
	var newCustomer dto.RegisterDTO

	// JSON verisini al
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	token, customerRes, err := st._service.Register(newCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"customer": customerRes,
		"token":    token,
	})
}
