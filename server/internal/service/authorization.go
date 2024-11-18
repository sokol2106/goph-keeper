// Package service предоставляет сервисы сисетмы
// Сервис управления авторизацией пользователей с использованием JWT (JSON Web Token) и хранения информации о пользователях в системе.

package service

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"sync/atomic"
)

// Authorization хранит информацию о пользователях системы и их текущем состоянии авторизации.
type Authorization struct {
	Users      sync.Map // ключ — login, значение — UserInfo
	CountUsers int64    // количество пользователей
}

// UserInfo представляет данные о пользователе.
type UserInfo struct {
	UserKey       string
	EncryptionKey string
}

// NewAuthorization создает новый объект Authorization и возвращает указатель на него.
func NewAuthorization() *Authorization {
	return &Authorization{
		CountUsers: 0,
	}
}

// NewUserToken создает новый JWT-токен для нового пользователя.
func (ath *Authorization) NewUserToken(login, userKey, encrKey string) (string, error) {
	ath.Users.Store(login, UserInfo{
		UserKey:       userKey,
		EncryptionKey: encrKey,
	})

	atomic.AddInt64(&ath.CountUsers, 1)
	token, err := NewToken(userKey, encrKey)
	if err != nil {
		return "", err
	}
	return token, err
}

// SetCurrentUserID в контексте запроса сохраняет текущего авторизованного пользователя.
func SetCurrentUserID(ctx context.Context, userKey uuid.UUID) context.Context {
	return context.WithValue(ctx, "currentUserKey", userKey)
}

// GetCurrentUserID извлекает текущего авторизованного пользователя из контекста запроса.
func GetCurrentUserID(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value("currentUserKey").(uuid.UUID)
	return userID, ok
}
