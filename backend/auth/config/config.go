package config

import (
	"log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Env string
	}
	Fiber struct {
		Host string
		Port string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if appConfig == nil {
		appConfig = &AppConfig{}

		initFiber(appConfig)
		initApp(appConfig)
	}
	return &AppConfig{}
}
