package main

import (
	"first-proj/appconfig"
	"first-proj/domain"
	"first-proj/services/connections"
	"first-proj/services/postgres"
	"first-proj/transport/httpt"
	"os"
	"os/signal"
	"syscall"

	"fmt"
)

var config = appconfig.MustLoad()

type DependencyContainer struct {
	noteService domain.NoteService
	httpApi     httpt.HttpApi
}

func main() {
	logger := appconfig.GetLogger()
	noteService := postgres.NewPostgres(
		connections.NewPool(
			int32(config.Storage.MaxConns),
			int32(config.Storage.MinConns),
			config.Storage.Url,
		),
	)
	di := DependencyContainer{
		noteService: noteService,
		httpApi:     httpt.NewHttpApiHandlers(noteService),
	}

	config := appconfig.MustLoad()
	fmt.Println(config)
	server := httpt.NewServer(config.Address, di.httpApi)
	done := make(chan struct{})
	server.Start(done)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
	logger.Info("Received Interrupt signal, shutting down...")

	close(done)

}
