package handlers

import (
	"io"
	"net/http"
)

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

func (h *Handlers) LogoutUser(w http.ResponseWriter, r *http.Request) {
	return
}
