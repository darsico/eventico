package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
    "email": email,
    "userId": userId,
    "exp": time.Now().Add(time.Hour * 2).Unix(),
  })
  return token.SignedString([]byte(SECRET_KEY))
}