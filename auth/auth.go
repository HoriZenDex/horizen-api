package auth

import (
	"os"
	"time"

	"horizen-api/models"

	"github.com/golang-jwt/jwt"
)

// GenerateJWT creates a JWT token for a given user.
func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix()
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
