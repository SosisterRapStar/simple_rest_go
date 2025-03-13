package main

import (
	"context"
	"first-proj/appconfig"
	"first-proj/domain"
	"first-proj/services/connections"
	"first-proj/services/postgres"
	"first-proj/transport/httpt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fmt"
)

var config = appconfig.MustLoad()

type DependencyContainer struct {
	noteService domain.NoteService
}

func main() {
	// store opened conns to close them after
	var openedConnections = []connections.Connection{}

	logger := appconfig.GetLogger()
	postgresConnection := &connections.PostgresConn{}
	if err := postgresConnection.Open(*config); err != nil {
		log.Fatal("Error opening connection to service")
	}
	openedConnections = append(openedConnections, postgresConnection)
	postgresService := postgres.NewPostgres(
		postgresConnection.Pool,
	)
	di := DependencyContainer{
		noteService: postgresService,
	}

	config := appconfig.MustLoad()
	fmt.Println(config)
	server := httpt.NewServer(config.Address, httpt.NewHttpApiHandlers(di.noteService))

	// starts the server
	server.Start()

	// waiting for SIGINT/SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	logger.Info("Received Interrupt signal, started to shutdown gracefully")
	gshutCtx, gshutClose := context.WithTimeout(context.Background(), 10*time.Second)
	defer gshutClose()

	// stopping the server gracefully
	server.Stop(gshutCtx)
	// closing other connections
	for _, conn := range openedConnections {
		closeCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		conn.Close(closeCtx)
	}
	logger.Info("App was stopped")
}
