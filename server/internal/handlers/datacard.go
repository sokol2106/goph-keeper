package handlers

import (
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"server/internal/service"
)

// CreateDataCard обрабатывает HTTP-запрос на добавление данных карты.
// Этот обработчик получает данные карты в теле запроса, проверяет авторизацию пользователя,
// а затем вставляет данные карты в систему. В случае успешной операции возвращает статус 201 (Created).
// В случае ошибки возвращает соответствующий HTTP-статус.
func (h *Handlers) CreateDataCard(w http.ResponseWriter, r *http.Request) {
	handlerStatus := http.StatusCreated
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handlerStatus = h.handlerError(err)
		if handlerStatus == http.StatusBadRequest {
			w.WriteHeader(handlerStatus)
			return
		}
	}
	userID, ok := service.GetCurrentUserID(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	resultBody, err := h.gophKeeper.InsertDataCard(body, userID)

	if err != nil {
		handlerStatus = h.handlerError(err)
		if handlerStatus == http.StatusBadRequest {
			w.WriteHeader(handlerStatus)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(handlerStatus)
	w.Write(resultBody)
}

// GetDataCard обрабатывает HTTP-запрос на получение данных карты по UUID.
// Этот обработчик извлекает UUID из параметров URL, проверяет авторизацию пользователя,
// а затем возвращает данные карты для указанного ключа. В случае успешной операции возвращает статус 200 (OK).
// В случае ошибки возвращается соответствующий HTTP-статус.
func (h *Handlers) GetDataCard(w http.ResponseWriter, r *http.Request) {
	handlerStatus := http.StatusOK
	key := chi.URLParam(r, "uuid")
	userID, ok := service.GetCurrentUserID(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	resultBody, err := h.gophKeeper.SelectDataCard(key, userID)

	if err != nil {
		handlerStatus = h.handlerError(err)
		if handlerStatus == http.StatusBadRequest {
			w.WriteHeader(handlerStatus)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(handlerStatus)
	w.Write(resultBody)
}

// DeleteDataCard обрабатывает HTTP-запрос на удаление данных карты по UUID.
// Этот обработчик извлекает UUID из параметров URL, проверяет авторизацию пользователя,
// а затем удаляет данные карты из системы. В случае успешной операции возвращает статус 200 (OK).
// В случае ошибки возвращается соответствующий HTTP-статус.
func (h *Handlers) DeleteDataCard(w http.ResponseWriter, r *http.Request) {
	handlerStatus := http.StatusOK
	key := chi.URLParam(r, "uuid")
	userID, ok := service.GetCurrentUserID(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := h.gophKeeper.DeleteDataCard(key, userID)

	if err != nil {
		handlerStatus = h.handlerError(err)
		if handlerStatus == http.StatusBadRequest {
			w.WriteHeader(handlerStatus)
			return
		}
	}

	w.WriteHeader(handlerStatus)
}
