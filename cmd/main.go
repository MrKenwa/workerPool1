package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"workerPool1/internal/config"
	"workerPool1/internal/logger"
	"workerPool1/internal/server"
)

func main() {
	//runtime.SetBlockProfileRate(1)
	//runtime.SetMutexProfileFraction(1)

	//f, _ := os.Create("cpu.out")
	//defer f.Close()
	//
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	log.SetOutput(os.Stdout)
	log.Print("Loading config...")
	cfg := config.Load()
	log.Print("Successful loaded config")

	log.Println("Try to init logger")
	logg, err := logger.New(cfg.Environment)
	if err != nil {
		log.Fatalf("Cannot initialize logger: %v", err)
	}
	logg.Info("Init logger successful")

	app := server.New(cfg, logg)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	if err := app.Start(); err != nil {
		log.Fatalf("error while starting server: %v", err)
	}
}
