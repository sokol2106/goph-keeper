package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Длительность жизни JWT-токена (24 часа).
const tokenEXP = time.Hour * 24

// Секретный ключ для подписи JWT-токенов.
const secretKey = "supersecret"

// Token описывает структуру JWT-токена с полем UserID.
type Token struct {
	jwt.RegisteredClaims
	UserKey       string `json:"user_key"`
	EncryptionKey string `json:"encryption_key"`
}

// ReadToken проверяет валидность JWT-токена и возвращает UserID из токена.
// Возвращает UserID и ошибку, если токен недействителен.
func ReadToken(cookValue string) (Token, error) {
	token := &Token{}

	res, err := jwt.ParseWithClaims(cookValue, token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return Token{}, err
	}

	if !res.Valid {
		return Token{}, errors.New("Token is not valid")
	}

	return *token, nil
}
