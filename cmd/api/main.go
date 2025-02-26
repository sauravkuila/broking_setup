package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sauravkuila/broking_setup/pkg/config"
	"github.com/sauravkuila/broking_setup/pkg/logger"
	"github.com/sauravkuila/broking_setup/pkg/server"
)

//package to start REST API server

func main() {
	var (
		environment string
	)
	host := os.Getenv("SERVER_HOST")
	if host != "" {
		environment = "server"
	} else {
		if len(os.Args) == 2 {
			environment = os.Args[1] // developer custom file
		} else {
			environment = "local"
		}
	}
	config.Load(environment)

	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server, err:", err)
		os.Exit(1)
	}
	addShutdownHook()
}

// adds and listens to any interrupt signal
//
//	the method should be called at the end of the main function as it blocks execution
//		internally it
//		-	closes redis connections
//		-	shuts down the http server
//		-	closes database connections
//	waits on syscall.SIGINT, syscall.SIGTERM, os.Interrupt
func addShutdownHook() {
	// when receive interruption from system shutdown server and scheduler
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	logger.Log().Info("Quit/Interrupt signal detected. Gracefully closing connections")
	//shutdown server
	server.ShutdownRouter()
	server.CloseDatabase()

	ctx := context.Background()

	logger.Log(ctx).Info(fmt.Sprintf("All done! Wrapping up here for PID: %d", os.Getpid()))
}
