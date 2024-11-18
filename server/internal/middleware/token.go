package middleware

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"server/internal/service"
)

// TokenResponseRequest является middleware-обработчиком, который проверяет наличие куки с токеном "user".
// Если куки не существует или токен недействителен, создает новый токен и устанавливает его в куки.
// Если токен существует и действителен, проверяет пользователя и продолжает выполнение запроса.
// В случае ошибки возвращает соответствующий HTTP-статус.
func TokenResponseRequest(gophKeeper *service.GophKeeper, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationPaths := []string{
			"/api/authorization",
			"/api/register",
		}

		// Пропускаем авторизацию
		for _, path := range authorizationPaths {
			if r.URL.Path == path {
				handler.ServeHTTP(w, r)
				return
			}
		}

		cookie, err := r.Cookie("user")
		// не существует или она не проходит проверку подлинности
		if err != nil {
			log.Printf("error handling request: %v, status: %d", err, http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
		}

		userKey, err := service.ReadToken(cookie.Value)
		if err != nil {
			log.Printf("error handling request: %v, status: %d", err, http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userKeyUUID, err := uuid.Parse(userKey)
		if err != nil {
			log.Printf("error handling request: %v, status: %d", err, http.StatusInternalServerError)
			return
		}

		ctx := service.SetCurrentUserID(r.Context(), userKeyUUID)

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
