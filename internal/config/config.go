package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server struct {
		Host string
	}
	Queue struct {
		Workers     int
		QueueSize   int
		BaseBackoff int
	}
}

const (
	defaultWorkersCount = 4
	defaultQueueSize    = 64
)

func Load() *Config {
	serverHost := os.Getenv("SERVER_HOST")
	workers := getEnvInt("WORKERS", defaultWorkersCount)
	qsize := getEnvInt("QUEUE_SIZE", defaultQueueSize)
	return &Config{
		Server: struct{ Host string }{
			Host: serverHost,
		},
		Queue: struct {
			Workers     int
			QueueSize   int
			BaseBackoff int
		}{
			Workers:     workers,
			QueueSize:   qsize,
			BaseBackoff: 100,
		},
	}
}

func getEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}
