package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set jwtKey from environment variable
	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}

// Claims defines JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// JWTAuthMiddleware verifies the token
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Extract token from Authorization header
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the username in the context for the next handlers
		c.Set("username", claims.Username)
		c.Next()
	}
}
