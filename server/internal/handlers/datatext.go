package handlers

import (
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"server/internal/service"
)

// CreateDataText обрабатывает HTTP-запрос на добавление текстовых данных.
// Этот обработчик получает данные в теле запроса, проверяет авторизацию пользователя,
// а затем вставляет текстовые данные в систему. В случае успешной операции возвращает статус 201 (Created).
// В случае ошибки возвращает соответствующий HTTP-статус.
func (h *Handlers) CreateDataText(w http.ResponseWriter, r *http.Request) {
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

	resultBody, err := h.gophKeeper.InsertDataText(body, userID)

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

// GetDataText обрабатывает HTTP-запрос на получение текстовых данных по UUID.
// Этот обработчик извлекает UUID из параметров URL, проверяет авторизацию пользователя,
// а затем возвращает текстовые данные для указанного ключа. В случае успешной операции возвращает статус 200 (OK).
// В случае ошибки возвращается соответствующий HTTP-статус.
func (h *Handlers) GetDataText(w http.ResponseWriter, r *http.Request) {
	handlerStatus := http.StatusOK
	key := chi.URLParam(r, "uuid")
	userID, ok := service.GetCurrentUserID(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	resultBody, err := h.gophKeeper.SelectDataText(key, userID)

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

// DeleteDataText обрабатывает HTTP-запрос на удаление текстовых данных по UUID.
// Этот обработчик извлекает UUID из параметров URL, проверяет авторизацию пользователя,
// а затем удаляет текстовые данные из системы. В случае успешной операции возвращает статус 200 (OK).
// В случае ошибки возвращается соответствующий HTTP-статус.
func (h *Handlers) DeleteDataText(w http.ResponseWriter, r *http.Request) {
	handlerStatus := http.StatusOK
	key := chi.URLParam(r, "uuid")
	userID, ok := service.GetCurrentUserID(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := h.gophKeeper.DeleteDataText(key, userID)

	if err != nil {
		handlerStatus = h.handlerError(err)
		if handlerStatus == http.StatusBadRequest {
			w.WriteHeader(handlerStatus)
			return
		}
	}

	w.WriteHeader(handlerStatus)
}
