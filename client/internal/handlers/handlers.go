package handlers

import (
	"client/internal/config"
	"client/internal/service"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

// Handlers представляет собой структуру, содержащую сервисы для обработки URL и авторизации.
type Handlers struct {
	gophKeeper *service.GophKeeperClient // Сервис сокращения URL
	cobra      *cobra.Command
	cnf        config.Config
	client     *http.Client
}

func NewHandlers(srv *service.GophKeeperClient, cnf config.Config) *Handlers {
	return &Handlers{
		gophKeeper: srv,
		cobra: &cobra.Command{
			Use:   "app",
			Short: "GophKeeper приложение",
		},
		cnf: cnf,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (h *Handlers) Run() error {
	h.cobra.AddCommand(
		h.RegisterUser(),
		h.AuthorizationUser(),
		h.CreateDataText(),
		h.GetDataText(),
		h.DeleteDataText(),
		h.CreateDataCard(),
		h.GetDataCard(),
		h.DeleteDataCard(),
		h.CreateDataBinary(),
		h.GetDataBinary(),
		h.DeleteDataBinary(),
	)

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
