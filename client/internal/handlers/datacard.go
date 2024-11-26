package handlers

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"net/url"
)

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
			if cardNumber == "" {
				fmt.Println("Укажите текст для отправки с помощью флага --text.")
				return
			}

			fmt.Println(cardNumber, cardholderName, expirationDate, cvvHash)
		},
	}

	// Добавляем флаги для команды
	cmd.Flags().StringVar(&cardNumber, "number", "", "Номер кредитной карты")
	cmd.Flags().StringVar(&cardholderName, "name", "", "Имя держателя карты")
	cmd.Flags().StringVar(&expirationDate, "date", "", "Дата истечения карты (MM/YY)")
	cmd.Flags().StringVar(&cvvHash, "cvv", "", "CVV")

	// Устанавливаем флаги как обязательные
	cmd.MarkFlagRequired("number")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("date")
	cmd.MarkFlagRequired("cvv")

	return cmd
}

func (h *Handlers) GetDataCard() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "getCard",
		Short: "Запрос карты",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.Flags().StringVar(&id, "key", "", "UUID данных")
	cmd.MarkFlagRequired("key")
	return cmd
}

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

	cmd.Flags().StringVar(&id, "key", "", "UUID данных")
	cmd.MarkFlagRequired("key")
	return cmd
}
