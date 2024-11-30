package handlers

import (
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"server/internal/service"
)

// CreateDataBinary обрабатывает HTTP-запрос на добавление бинарных данных.
// Этот обработчик получает бинарные данные в теле запроса, проверяет авторизацию пользователя,
// а затем вставляет данные в систему. В случае успешной операции возвращает статус 201 (Created).
// В случае ошибки возвращается соответствующий HTTP-статус.
func (h *Handlers) CreateDataBinary(w http.ResponseWriter, r *http.Request) {
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

	resultBody, err := h.gophKeeper.InsertDataBinary(body, userID)

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

// GetDataBinary обрабатывает HTTP-запрос на получение бинарных данных по UUID.
// Этот обработчик извлекает UUID из параметров URL, проверяет авторизацию пользователя,
// а затем возвращает бинарные данные для указанного ключа. В случае успешной операции возвращает статус 200 (OK).
// В случае ошибки возвращается соответствующий HTTP-статус.
func (h *Handlers) GetDataBinary(w http.ResponseWriter, r *http.Request) {
	handlerStatus := http.StatusOK
	key := chi.URLParam(r, "uuid")
	userID, ok := service.GetCurrentUserID(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	resultBody, err := h.gophKeeper.SelectDataBinary(key, userID)

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

// DeleteDataBinary обрабатывает HTTP-запрос на удаление бинарных данных по UUID.
// Этот обработчик извлекает UUID из параметров URL, проверяет авторизацию пользователя,
// а затем удаляет данные из системы. В случае успешной операции возвращает статус 200 (OK).
// В случае ошибки возвращается соответствующий HTTP-статус.
func (h *Handlers) DeleteDataBinary(w http.ResponseWriter, r *http.Request) {
	handlerStatus := http.StatusOK
	key := chi.URLParam(r, "uuid")
	userID, ok := service.GetCurrentUserID(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := h.gophKeeper.DeleteDataBinary(key, userID)

	if err != nil {
		handlerStatus = h.handlerError(err)
		if handlerStatus == http.StatusBadRequest {
			w.WriteHeader(handlerStatus)
			return
		}
	}

	w.WriteHeader(handlerStatus)
}
