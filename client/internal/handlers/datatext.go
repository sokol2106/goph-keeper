package handlers

import (
	"bytes"
	"client/internal/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// CreateDataText создает команду для добавления текстовых данных.
// Пример:
// addText --text "Hello, World!"
// Ответ:
// dataTextKey: "abcd1234"
func (h *Handlers) CreateDataText() *cobra.Command {
	var text string
	cmd := &cobra.Command{
		Use:   "addText",
		Short: "Добавление текстовых данных",
		Run: func(cmd *cobra.Command, args []string) {
			codText, err := h.gophKeeper.CreateText(text)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			data := model.DataText{Data: codText}

			reqBody, _ := json.Marshal(data)
			req, err := http.NewRequest(http.MethodPost, h.cnf.Listen+"/api/data/text", strings.NewReader(string(reqBody)))
			if err != nil {
				log.Printf("%v", err)
				return
			}
			req.AddCookie(h.gophKeeper.GetCookie())

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			textResponse := model.DataTextResponse{}
			err = json.NewDecoder(resp.Body).Decode(&textResponse)
			resp.Body.Close()
			if err != nil {
				log.Printf("%v", err)
				return
			}

			fmt.Println(textResponse.DataTextKey)
		},
	}

	cmd.Flags().StringVarP(&text, "text", "t", "", "Текст для отправки")
	cmd.MarkFlagRequired("text")
	return cmd
}

// GetDataText создает команду для запроса текстовых данных по их UUID.
// При выполнении команды отправляется запрос на сервер с указанным UUID.
// Сервер возвращает текст, который выводится на экран.
// Пример:
// getText --key abcd1234
// Ответ:
// "Hello, World!"
func (h *Handlers) GetDataText() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "getText",
		Short: "Запрос текстовых данных",
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := uuid.Parse(id); err != nil {
				log.Printf("UUID Parser: %v", err)
				return
			}

			req, err := http.NewRequest(http.MethodGet, h.cnf.Listen+"/api/data/text/"+id, bytes.NewBuffer(nil))
			if err != nil {
				log.Printf("%v", err)
				return
			}
			req.AddCookie(h.gophKeeper.GetCookie())

			resp, err := h.client.Do(req)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			body, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Printf("Ошибка: сервер вернул ошибочный статус: %d %s", resp.StatusCode, resp.Status)
				return
			}

			result, err := h.gophKeeper.GetText(body)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			fmt.Println(result)
		},
	}

	cmd.Flags().StringVarP(&id, "key", "k", "", "UUID данных")
	cmd.MarkFlagRequired("key")
	return cmd
}

// DeleteDataText создает команду для удаления текстовых данных по их UUID.
// При выполнении команды отправляется запрос на сервер с указанным UUID для удаления данных.
// После успешного удаления выводится сообщение о результате операции.
// Пример:
// delText --key abcd1234
// Ответ:
// "Данные удалены"
func (h *Handlers) DeleteDataText() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delText",
		Short: "Удаление текстовых данных",
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := uuid.Parse(id); err != nil {
				log.Printf("UUID Parser: %v", err)
				return
			}

			reqURL := fmt.Sprintf("%s/api/data/text/%s", h.cnf.Listen, url.PathEscape(id))
			req, err := http.NewRequest(http.MethodDelete, reqURL, bytes.NewBuffer(nil))
			if err != nil {
				log.Printf("%v", err)
				return
			}
			req.AddCookie(h.gophKeeper.GetCookie())

			resp, err := h.client.Do(req)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				fmt.Println("Данные удалены")
			} else {
				log.Printf("Ошибка удаления: HTTP %d - %s\n", resp.StatusCode, resp.Status)
			}
		},
	}

	cmd.Flags().StringVarP(&id, "key", "k", "", "UUID данных")
	cmd.MarkFlagRequired("key")
	return cmd
}
