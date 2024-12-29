package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Memeriksa format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// Mendapatkan SECRET dari environment
		secret := os.Getenv("SECRET")
		if secret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing SECRET in environment variables"})
			c.Abort()
			return
		}

		// Memvalidasi token
		token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
			// Memeriksa metode signing
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Menyimpan klaim di context untuk digunakan di handler
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Mengonversi klaim "id" ke tipe yang diinginkan (misalnya uint)
			id, ok := claims["id"].(float64) // Klaim ID dalam JWT biasanya bertipe float64
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID type in token"})
				c.Abort()
				return
			}

			// Menyimpan ID dan email ke context, pastikan ID adalah tipe yang benar (uint)
			c.Set("user_id", uint(id)) // Mengonversi float64 ke uint
			c.Set("email", claims["email"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}
