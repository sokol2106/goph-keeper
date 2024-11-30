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

// CreateDataCard создает команду для добавления данных кредитной карты.
// При выполнении команды отправляется запрос с данными карты на сервер.
// После успешной отправки, сервер возвращает уникальный ключ для этих данных,
// который выводится на экран.
// Пример:
// addCard --number "1234567812345678" --name "John Doe" --date "12/25" --cvv "123"
// Ответ:
// dataCreditCardKey: "abcd1234"
func (h *Handlers) CreateDataCard() *cobra.Command {
	var (
		cardNumber     string
		cardholderName string
		expirationDate string
		cvvHash        string
	)
	cmd := &cobra.Command{
		Use:   "addCard",
		Short: "Добавление карты",
		Run: func(cmd *cobra.Command, args []string) {
			mod := model.DataCreditCardResponse{
				CardNumber:     cardNumber,
				CardholderName: cardholderName,
				ExpirationDate: expirationDate,
				CVVHash:        cvvHash,
			}

			modEnc, err := h.gophKeeper.CreateCreditCard(mod)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			reqBody, _ := json.Marshal(modEnc)
			req, err := http.NewRequest(http.MethodPost, h.cnf.Listen+"/api/data/card", strings.NewReader(string(reqBody)))
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

			textResponse := model.DataCreditCardResponse{}
			err = json.NewDecoder(resp.Body).Decode(&textResponse)
			resp.Body.Close()
			if err != nil {
				log.Printf("%v", err)
				return
			}

			fmt.Println(textResponse.DataCreditCardKey)
		},
	}

	// Добавляем флаги для команды
	cmd.Flags().StringVarP(&cardNumber, "number", "q", "", "Номер кредитной карты")
	cmd.Flags().StringVarP(&cardholderName, "name", "n", "", "Имя держателя карты")
	cmd.Flags().StringVarP(&expirationDate, "date", "d", "", "Дата истечения карты (MM/YY)")
	cmd.Flags().StringVarP(&cvvHash, "cvv", "c", "", "CVV")

	// Устанавливаем флаги как обязательные
	cmd.MarkFlagRequired("number")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("date")
	cmd.MarkFlagRequired("cvv")

	return cmd
}

// GetDataCard создает команду для запроса данных кредитной карты по UUID.
// При выполнении команды отправляется запрос на сервер с указанным UUID.
// Сервер возвращает информацию о карте, которая выводится на экран.
// Пример:
// getCard --key "abcd1234"
// Ответ:
// "abcd1234"
// "1234567812345678"
// "John Doe"
// "12/25"
// "123"
// "2024-11-30T12:34:56Z"
func (h *Handlers) GetDataCard() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "getCard",
		Short: "Запрос карты",
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := uuid.Parse(id); err != nil {
				log.Printf("UUID Parser: %v", err)
				return
			}

			req, err := http.NewRequest(http.MethodGet, h.cnf.Listen+"/api/data/card/"+id, bytes.NewBuffer(nil))
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

			result, err := h.gophKeeper.GetCreditCard(body)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			fmt.Println(result.DataCreditCardKey)
			fmt.Println(result.CardNumber)
			fmt.Println(result.CardholderName)
			fmt.Println(result.ExpirationDate)
			fmt.Println(result.CVVHash)
			fmt.Println(result.CreatedAt)
		},
	}

	cmd.Flags().StringVarP(&id, "key", "k", "", "UUID данных")
	cmd.MarkFlagRequired("key")
	return cmd
}

// DeleteDataCard создает команду для удаления данных кредитной карты по UUID.
// При выполнении команды отправляется запрос на сервер с указанным UUID для удаления данных.
// После успешного удаления выводится сообщение о результате операции.
// Пример:
// delCard --key "abcd1234"
// Ответ:
// "Данные удалены"
func (h *Handlers) DeleteDataCard() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "delCard",
		Short: "Удаление карты",
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := uuid.Parse(id); err != nil {
				log.Printf("UUID Parser: %v", err)
				return
			}

			reqURL := fmt.Sprintf("%s/api/data/card/%s", h.cnf.Listen, url.PathEscape(id))
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
