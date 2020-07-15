package utils

import (
	"time"
	"os"

	"github.com/dgrijalva/jwt-go"

)

//JWTSecret to set secret key
var JWTSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))

//GenerateJWT to generate token
func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
