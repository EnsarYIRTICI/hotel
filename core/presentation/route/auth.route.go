package route

import (
	"go-basic/application/service"
	"go-basic/presentation/dto"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Tüm Auth route'larını kaydeder
func RegisterAuthRoutes(router *gin.Engine) {

	authService := service.NewAuthService()

	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("register", func(c *gin.Context) { Register(c, authService) })
		authGroup.POST("login", func(c *gin.Context) { Login(c, authService) })
	}
}

// Login, kullanıcı girişini başlatır
func Login(c *gin.Context, authService service.AuthService) {
	var loginData dto.LoginDTO

	// JSON verisini al
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri formatı"})
		return
	}

	// AuthService ile işlemi gerçekleştir
	token, customerRes, err := authService.Login(loginData)
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
func Register(c *gin.Context, authService service.AuthService) {
	var newCustomer dto.RegisterDTO

	// JSON verisini al
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	token, customerRes, err := authService.Register(newCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"customer": customerRes,
		"token":    token,
	})
}
