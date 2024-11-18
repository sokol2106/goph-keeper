package handlers

import (
	"client/internal/service"
	"github.com/spf13/cobra"
)

// Handlers представляет собой структуру, содержащую сервисы для обработки URL и авторизации.
type Handlers struct {
	gophKeeper *service.GophKeeperClient // Сервис сокращения URL
	cobra      *cobra.Command
}

func NewHandlers(srv *service.GophKeeperClient) *Handlers {
	return &Handlers{
		gophKeeper: srv,
		cobra: &cobra.Command{
			Use:   "app",
			Short: "GophKeeper приложение",
		},
	}
}

func (h *Handlers) Run() error {
	h.cobra.AddCommand(h.RegisterUser())
	if err := h.cobra.Execute(); err != nil {
		return err
	}
	return nil
}

func (h *Handlers) Execute() error {
	return h.cobra.Execute()
}

func (h *Handlers) SetArgs(args []string) {
	h.cobra.SetArgs(args)
}
