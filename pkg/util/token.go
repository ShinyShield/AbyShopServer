package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func CreateToken(email string, pwd string) (string, error) {
	claim := jwt.MapClaims{
		"id":       email,
		"username": pwd,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(viper.GetString("token.secret")))
}
