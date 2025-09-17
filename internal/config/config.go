package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
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
	serverHost := os.Getenv("SERVER_HOST")
	workers := getEnvInt("WORKERS", defaultWorkersCount)
	qsize := getEnvInt("QUEUE_SIZE", defaultQueueSize)
	return &Config{
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
