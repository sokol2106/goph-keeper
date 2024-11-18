package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"server/internal/config"
	"server/internal/handlers"
	"server/internal/server"
	"server/internal/service"
	"server/internal/storage"
	"syscall"
)

func Run(cnf *config.Config) {
	objStorage := storage.NewPostgresql(*cnf)
	err := objStorage.Connect()
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
	objService := service.NewGophKeeper(objStorage)
	objHandler := handlers.NewHandlers(&objService)
	objServer := server.NewServer(server.Router(objHandler), cnf.Listen)

	idleConnsClosed := make(chan struct{})
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-stop
		if err := objServer.Stop(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}

		if err := objStorage.Close(); err != nil {
			log.Printf("Object storage close: %v", err)
		}

		log.Println("Signal shutdown")
		close(idleConnsClosed)
	}()

	err = objServer.Start()
	if err != nil {
		log.Printf("Starting server error: %s", err)
	}

	<-idleConnsClosed
	fmt.Println("Server Shutdown gracefully")
}
