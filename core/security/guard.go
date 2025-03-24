package security

import (
	"strings"

	"go-basic/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Guard fonksiyonu, JWT doğrulama işlemini yapar
func Guard(exemptRoutes []string, logger util.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Rota kontrolü: Eğer istek yapılan rota exemptRoutes içinde varsa, Guard middleware'ini atla

		logger.Log("ÇALIŞTI")

		for _, route := range exemptRoutes {

			if c.FullPath() == route {
				c.Next() // Rota listede varsa, middleware'i atla
				return
			}
		}

		// Authorization header'ını al
		tokenString := c.GetHeader("Authorization")

		// Eğer Authorization header'ı yoksa
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No token provided"})
			c.Abort()
			return
		}

		// Token'ı "Bearer <token>" formatında doğrula
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token format"})
			c.Abort()
			return
		}

		// "Bearer " kısmını çıkar ve sadece token'ı al
		tokenString = tokenString[7:]

		// Token'ı doğrulamak için jwt.ValidateJWT fonksiyonunu çağır
		claims, err := ValidateJWT(tokenString)

		// Eğer token geçersizse ya da çözümleme hatası varsa
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid or expired token"})
			c.Abort()
			return
		}

		// Token geçerli ise, context'e kullanıcı bilgisi ekleyebiliriz
		c.Set("identity_number", claims.IdentityNumber)

		// Sonraki middleware'e veya handler'a geçiş
		c.Next()
	}
}
