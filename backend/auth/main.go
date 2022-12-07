package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alfinka/backend/auth/app/server"
)

func main() {
	app := server.Server()

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
