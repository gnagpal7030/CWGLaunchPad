package auth

import (
	"CWDLaunchPad/constants"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string) (string, error) {
	secretKey := []byte(os.Getenv(constants.SecretKey))
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	return token.SignedString(secretKey)
}
