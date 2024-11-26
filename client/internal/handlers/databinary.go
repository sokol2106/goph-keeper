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

func (h *Handlers) CreateDataBinary() *cobra.Command {
	var filename string
	cmd := &cobra.Command{
		Use:   "addBinary",
		Short: "Добавление бинарных данных",
		Run: func(cmd *cobra.Command, args []string) {
			if filename == "" {
				fmt.Println("Укажите текст для отправки с помощью флага --text.")
				return
			}

			fmt.Println(filename)
		},
	}

	// Добавляем флаги для команды
	cmd.Flags().StringVarP(&filename, "filename", "f", "", "Файл")

	// Устанавливаем флаги как обязательные
	cmd.MarkFlagRequired("filename")

	return cmd
}

func (h *Handlers) GetDataBinary() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "getBinary",
		Short: "Запрос бинарных данных",
		Run: func(cmd *cobra.Command, args []string) {
			if id == "" {
				fmt.Println("Укажите параметр --uuid для выполнения запроса.")
				return
			}
			fmt.Println("Goood.")
		},
	}

	cmd.Flags().StringVar(&id, "key", "", "UUID данных")
	cmd.MarkFlagRequired("key")
	return cmd
}

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

	cmd.Flags().StringVar(&id, "key", "", "UUID данных")
	cmd.MarkFlagRequired("key")
	return cmd
}
