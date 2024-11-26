package handlers

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"net/url"
)

func (h *Handlers) CreateDataText() *cobra.Command {
	var text string
	cmd := &cobra.Command{
		Use:   "addText",
		Short: "Добавление текстовых данных",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Goood.")
		},
	}

	cmd.Flags().StringVarP(&text, "text", "t", "", "Текст для отправки")
	cmd.MarkFlagRequired("text")
	return cmd
}

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

	cmd.Flags().StringVar(&id, "key", "", "UUID данных")
	cmd.MarkFlagRequired("key")
	return cmd
}

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

	cmd.Flags().StringVar(&id, "key", "", "UUID данных")
	cmd.MarkFlagRequired("key")
	return cmd
}
