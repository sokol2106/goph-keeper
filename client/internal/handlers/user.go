package handlers

import (
	"bytes"
	"client/internal/model"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

// RegisterUser создает команду для регистрации пользователя.
// При выполнении команды запрашивает у пользователя логин и пароль,
// формирует запрос к серверу для регистрации пользователя.
// В случае успеха сохраняет cookie от сервера.
func (h *Handlers) RegisterUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reg",
		Short: "Регистрация пользователя",
		Run: func(cmd *cobra.Command, args []string) {
			var username, password string
			fmt.Print("Введите логи: ")
			fmt.Scanln(&username)
			fmt.Print("Введите пароль: ")
			fmt.Scanln(&password)

			data := model.User{
				Login:        username,
				PasswordHash: password,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			req, err := http.NewRequest(http.MethodPost, h.cnf.Listen+"/api/register", bytes.NewBuffer(jsonData))
			if err != nil {
				log.Printf("Ошибка при создании запроса: %v", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err := h.client.Do(req)
			if err != nil {
				log.Printf("Ошибка при отправке запроса: %v", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				log.Printf("Ошибка: сервер вернул ошибочный статус: %d %s", resp.StatusCode, resp.Status)
				return
			}

			h.gophKeeper.SetCookie(resp.Cookies()[0])
		},
	}

	return cmd
}

// AuthorizationUser создает команду для авторизации пользователя.
// При выполнении команды запрашивает у пользователя логин и пароль,
// формирует запрос к серверу для авторизации.
// В случае успеха сохраняет cookie от сервера.
func (h *Handlers) AuthorizationUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aut",
		Short: "Авторизация пользователя",
		Run: func(cmd *cobra.Command, args []string) {
			var username, password string
			fmt.Print("Введите логи: ")
			fmt.Scanln(&username)
			fmt.Print("Введите пароль: ")
			fmt.Scanln(&password)

			data := model.User{
				Login:        username,
				PasswordHash: password,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			req, err := http.NewRequest(http.MethodPost, h.cnf.Listen+"/api/authorization", bytes.NewBuffer(jsonData))
			if err != nil {
				log.Printf("Ошибка при создании запроса: %v", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err := h.client.Do(req)
			if err != nil {
				log.Printf("Ошибка при отправке запроса: %v", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				log.Printf("Ошибка: сервер вернул ошибочный статус: %d %s", resp.StatusCode, resp.Status)
				return
			}

			h.gophKeeper.SetCookie(resp.Cookies()[0])

		},
	}

	return cmd
}
