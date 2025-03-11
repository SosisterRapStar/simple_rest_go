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
}

func main() {
	logger := appconfig.GetLogger()
	postgresService := postgres.NewPostgres(
		connections.NewPool(
			int32(config.Storage.MaxConns),
			int32(config.Storage.MinConns),
			config.Storage.Url,
		),
	)
	di := DependencyContainer{
		noteService: postgresService,
	}

	config := appconfig.MustLoad()
	fmt.Println(config)
	server := httpt.NewServer(config.Address, httpt.NewHttpApiHandlers(di.noteService))
	done := make(chan struct{})
	server.Start(done)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
	logger.Info("Received Interrupt signal, shutting down...")

	close(done)

}
