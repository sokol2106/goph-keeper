package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
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

// NewToken создает и возвращает новый JWT-токен с указанным userID.
// Возвращает строку с токеном или ошибку.
func NewToken(userKey, encryptionKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Token{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenEXP)),
		},
		UserKey:       userKey,
		EncryptionKey: encryptionKey,
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ReadToken проверяет валидность JWT-токена и возвращает UserID из токена.
// Возвращает UserID и ошибку, если токен недействителен.
func ReadToken(cookValue string) (string, error) {
	token := &Token{}

	res, err := jwt.ParseWithClaims(cookValue, token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	if !res.Valid {
		return "", errors.New("Token is not valid")
	}

	return token.UserKey, nil
}

// GenerateEncryptionKey создает криптографически безопасный ключ шифрования.
// length — длина ключа в байтах.
func GenerateEncryptionKey() (string, error) {
	// Создаем слайс для хранения случайных байт.
	key := make([]byte, 32)

	// Заполняем слайс криптографически случайными байтами.
	_, err := rand.Read(key)
	if err != nil {
		return "", fmt.Errorf("failed to generate encryption key: %w", err)
	}

	// Кодируем ключ в base64 для удобной передачи и хранения.
	return base64.StdEncoding.EncodeToString(key), nil
}
