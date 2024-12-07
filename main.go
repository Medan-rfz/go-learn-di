package main

import (
	"flag"
	"fmt"
	"log"

	"test-di/internal/handler"
	"test-di/internal/repository"
	"test-di/internal/server"

	"go.uber.org/dig"
)

type appConfig struct {
	port      int
	dbConnStr string
}

// @title API
func main() {
	cfg := getAppConfig()

	container := initDI(cfg)
	err := container.Invoke(func(server *server.Server) {
		host := fmt.Sprintf(":%v", cfg.port)

		log.Printf("Server starting [%s]", host)

		if err := server.Run(host); err != nil {
			log.Fatalln(err)
		}
	})
	if err != nil {
		log.Fatalln("Error invoking server:", err)
	}
}

func initDI(cfg appConfig) *dig.Container {
	container := dig.New()

	err := container.Provide(func() string {
		return cfg.dbConnStr
	})
	if err != nil {
		log.Fatalln("Error providing connection string:", err)
	}

	err = container.Provide(repository.NewRepository, dig.As(new(handler.IRepo)))
	if err != nil {
		log.Fatalln("Error providing repository:", err)
	}

	err = container.Provide(handler.NewHandler, dig.As(new(server.IHandler)))
	if err != nil {
		log.Fatalln("Error providing handler:", err)
	}

	err = container.Provide(server.NewServer)
	if err != nil {
		log.Fatalln("Error providing server:", err)
	}

	return container
}

func getAppConfig() (cfg appConfig) {
	flag.IntVar(&cfg.port, "port", 8080, "The port used by the application")
	flag.StringVar(&cfg.dbConnStr, "db", "database_connection_string", "Database connecting string")
	flag.Parse()

	return
}
