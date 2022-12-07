package config

import (
	"log"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func initFiber(conf *AppConfig) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	switch {
	case port == "":
		log.Printf("Port is not set. Using default: %s", port)
		port = PORT
	case host == "":
		log.Printf("Host is not set. Using default: %s", HOST)
		host = HOST
	}

	conf.Fiber.Host = HOST
	conf.Fiber.Port = PORT
}
