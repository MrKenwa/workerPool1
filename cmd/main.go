package main

import (
	"log"
	"os"

	"workerPool1/internal/config"
	"workerPool1/internal/server"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Print("Loading config...")
	cfg := config.Load()
	log.Print("Successful loaded config")

	app := server.New(cfg)

	if err := app.Start(); err != nil {
		log.Fatalf("error while starting server: %v", err)
	}
}
