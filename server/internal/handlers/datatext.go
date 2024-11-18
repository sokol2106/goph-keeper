package handlers

import (
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"server/internal/service"
)

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

func (h *Handlers) GetDataText(w http.ResponseWriter, r *http.Request) {
	handlerStatus := http.StatusCreated
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
