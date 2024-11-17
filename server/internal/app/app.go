package app

import (
	"server/internal/config"
	"server/internal/service"
	"server/internal/storage"
)

func Run(cnf *config.Config) {

	str := storage.NewPostgresql("")
	service.NewGophKeeper(str)
}
