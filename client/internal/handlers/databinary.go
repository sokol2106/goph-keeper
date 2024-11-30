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

// CreateDataBinary создает команду для добавления бинарных данных на сервер.
// При выполнении команды отправляется запрос с бинарными данными (файл) на сервер.
// После успешной отправки, сервер возвращает уникальный ключ для этих данных,
// который выводится на экран.
// Пример:
// addBinary --name "file.txt" --path "/path/to/file.txt"
// Ответ:
// dataBinaryKey: "abcd1234"
func (h *Handlers) CreateDataBinary() *cobra.Command {
	var (
		filename string
		path     string
	)
	cmd := &cobra.Command{
		Use:   "addBinary",
		Short: "Добавление бинарных данных",
		Run: func(cmd *cobra.Command, args []string) {
			mod, err := h.gophKeeper.CreateBinary(filename, path)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			reqBody, _ := json.Marshal(mod)
			req, err := http.NewRequest(http.MethodPost, h.cnf.Listen+"/api/data/binary", strings.NewReader(string(reqBody)))
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

			textResponse := model.DataBinaryResponse{}
			err = json.NewDecoder(resp.Body).Decode(&textResponse)
			resp.Body.Close()
			if err != nil {
				log.Printf("%v", err)
				return
			}

			fmt.Println(textResponse.DataBinaryKey)
		},
	}

	// Добавляем флаги для команды
	cmd.Flags().StringVarP(&filename, "name", "n", "", "Наименование")
	cmd.Flags().StringVarP(&path, "path", "p", "", "Путь")

	// Устанавливаем флаги как обязательные
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("path")
	return cmd
}


// GetDataBinary создает команду для запроса бинарных данных по UUID.
// При выполнении команды отправляется запрос на сервер с указанным UUID.
// Сервер возвращает бинарные данные, и создаёт файл рядом с программой
// Пример:
// getBinary --key "abcd1234"
// Ответ:
// Путь к созданному файлу
func (h *Handlers) GetDataBinary() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "getBinary",
		Short: "Запрос бинарных данных",
		Run: func(cmd *cobra.Command, args []string) {

			if _, err := uuid.Parse(id); err != nil {
				log.Printf("UUID Parser: %v", err)
				return
			}

			req, err := http.NewRequest(http.MethodGet, h.cnf.Listen+"/api/data/binary/"+id, bytes.NewBuffer(nil))
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

			result, err := h.gophKeeper.GetBinary(body)
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


// DeleteDataBinary создает команду для удаления бинарных данных по UUID.
// При выполнении команды отправляется запрос на сервер с указанным UUID для удаления данных.
// После успешного удаления выводится сообщение о результате операции.
// Пример:
// delBinary --key "abcd1234"
// Ответ:
// "Данные удалены"
func (h *Handlers) DeleteDataBinary() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delBinary",
		Short: "Удаление бинарных данных",
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := uuid.Parse(id); err != nil {
				log.Printf("UUID Parser: %v", err)
				return
			}

			reqURL := fmt.Sprintf("%s/api/data/binary/%s", h.cnf.Listen, url.PathEscape(id))
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
