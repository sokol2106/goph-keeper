package handlers

import (
	"io"
	"net/http"
)

// RegisterUser обрабатывает HTTP-запрос на регистрацию нового пользователя.
// Этот обработчик получает данные для регистрации из тела запроса,
// вызывает метод для регистрации пользователя, а затем возвращает ответ с результатом.
// В случае успешной регистрации возвращается статус 201 (Created) и тело ответа с результатами.
// Также устанавливается cookie с токеном авторизации пользователя.
func (h *Handlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
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

	resultBody, token, err := h.gophKeeper.RegisterUser(string(body))

	if err != nil {
		handlerStatus = h.handlerError(err)
		if handlerStatus == http.StatusBadRequest {
			w.WriteHeader(handlerStatus)
			return
		}
	}

	newCookie := http.Cookie{Name: "user", Value: token}
	http.SetCookie(w, &newCookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(handlerStatus)
	w.Write(resultBody)
}

// AuthorizationUser обрабатывает HTTP-запрос на авторизацию пользователя.
// Этот обработчик получает данные для авторизации из тела запроса,
// вызывает метод для авторизации пользователя, а затем возвращает ответ с результатом.
// В случае успешной авторизации возвращается статус 201 (Created) и тело ответа с результатами.
// Также устанавливается cookie с токеном авторизации пользователя.
func (h *Handlers) AuthorizationUser(w http.ResponseWriter, r *http.Request) {
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

	resultBody, token, err := h.gophKeeper.AuthorizationUser(string(body))

	if err != nil {
		handlerStatus = h.handlerError(err)
		if handlerStatus == http.StatusBadRequest {
			w.WriteHeader(handlerStatus)
			return
		}
	}

	newCookie := http.Cookie{Name: "user", Value: token}
	http.SetCookie(w, &newCookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(handlerStatus)
	w.Write(resultBody)
}
