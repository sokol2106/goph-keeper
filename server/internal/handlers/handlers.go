package handlers

import (
	"server/internal/service"
)

// Handlers представляет собой структуру, содержащую сервисы для обработки URL и авторизации.
type Handlers struct {
	srvShortURL *service.GophKeeper // Сервис сокращения URL
}

// NewHandlers создает новый экземпляр Handlers с переданным сервисом сокращения URL.
func NewHandlers(srv *service.GophKeeper) *Handlers {
	return &Handlers{
		srvShortURL: srv,
	}
}
