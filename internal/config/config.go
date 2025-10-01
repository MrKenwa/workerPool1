package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Environment  string
	ServerConfig ServerConfig
	QueueConfig  QueueConfig
}

type ServerConfig struct {
	Host string
}

type QueueConfig struct {
	Workers     int
	QueueSize   int
	BaseBackoff time.Duration
}

const (
	defaultWorkersCount = 4
	defaultQueueSize    = 64
)

func Load() *Config {
	serverHost := func() string {
		res := os.Getenv("SERVER_HOST")
		if res != "" {
			return res
		}
		return "localhost:8080"
	}()
	workers := getEnvInt("WORKERS", defaultWorkersCount)
	qsize := getEnvInt("QUEUE_SIZE", defaultQueueSize)
	return &Config{
		Environment: os.Getenv("ENVIRONMENT"),
		ServerConfig: ServerConfig{
			Host: serverHost,
		},
		QueueConfig: QueueConfig{
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
