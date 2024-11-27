package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"server/internal/handlers"
	"server/internal/middleware"
)

func Router(h *handlers.Handlers) chi.Router {
	router := chi.NewRouter()

	// middleware
	//router.Use(middleware.СompressionResponseRequest)
	router.Use(middleware.LoggingResponseRequest)
	router.Use(func(handlerF http.Handler) http.Handler {
		return middleware.TokenResponseRequest(h.GetServiceGophKeeper(), handlerF)
	})

	// router
	// user
	router.Post("/api/register", http.HandlerFunc(h.RegisterUser))
	router.Post("/api/authorization", http.HandlerFunc(h.AuthorizationUser))

	// data
	// data text
	router.Post("/api/data/text", http.HandlerFunc(h.CreateDataText))
	router.Get("/api/data/text/{uuid}", http.HandlerFunc(h.GetDataText))
	router.Delete("/api/data/text/{uuid}", http.HandlerFunc(h.DeleteDataText))

	// data byte
	router.Post("/api/data/binary", http.HandlerFunc(h.CreateDataBinary))
	router.Get("/api/data/binary/{uuid}", http.HandlerFunc(h.GetDataBinary))
	router.Delete("/api/data/binary/{uuid}", http.HandlerFunc(h.DeleteDataBinary))

	// data card
	router.Post("/api/data/card", http.HandlerFunc(h.CreateDataCard))
	router.Get("/api/data/card/{uuid}", http.HandlerFunc(h.GetDataCard))
	router.Delete("/api/data/card/{uuid}", http.HandlerFunc(h.DeleteDataCard))

	return router
}
