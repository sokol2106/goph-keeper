package handlers

import (
	"log"
	"net/http"
	"server/internal/service"
)

// Handlers представляет собой структуру, содержащую сервисы для обработки URL и авторизации.
type Handlers struct {
	gophKeeper *service.GophKeeper // Сервис сокращения URL
}

// NewHandlers создает новый экземпляр Handlers с переданным сервисом сокращения URL.
func NewHandlers(srv *service.GophKeeper) *Handlers {
	return &Handlers{
		gophKeeper: srv,
	}
}

// handlerError обрабатывает ошибки и возвращает соответствующий код состояния HTTP.
// Следующие коды могут вернуться:
// - 400 Bad Request: для всех прочих ошибок.
// - 409 Conflict: если пытаетесь добавить уже существующий оригинальный URL.
// - 410 Gone: если URL был помечен как удаленный.
func (h *Handlers) handlerError(err error) int {
	statusCode := http.StatusBadRequest
	//if errors.Is(err, cerrors.ErrNewShortURL) {
	//	statusCode = http.StatusConflict
	//	}
	//if errors.Is(err, cerrors.ErrGetShortURLDelete) {
	//	statusCode = http.StatusGone
	//	}

	log.Printf("error handling request: %v, status: %d", err, statusCode)
	return statusCode
}

func (h *Handlers) GetServiceGophKeeper() *service.GophKeeper {
	return h.gophKeeper
}
