package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alfinka/backend/auth/app/server"
	"github.com/alfinka/backend/auth/config"
)

func main() {
	app := server.Server()
	conf := config.NewAppConfig()

	err := app.Listen(conf.Fiber.Host + ":" + conf.Fiber.Port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
