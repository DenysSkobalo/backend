// utils/jwt_utils.go

package utils

import (
	"crypto/rand"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey []byte

func init() {
	keyLength := 32
	key, err := GenerateRandomKey(keyLength)
	if err != nil {
		panic("Помилка генерації секретного ключа JWT")
	}
	JwtKey = key
}

// GenerateToken generates a JWT token for the specified user
func GenerateToken(username string) (string, error) {
	// create a label (payload) for the token
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 hours

	// create a token with the specified stamp and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRandomKey(length int) ([]byte, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}
