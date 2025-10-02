package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Environment  string `envconfig:"ENVIRONMENT"`
	ServerConfig ServerConfig
	QueueConfig  QueueConfig
}

type ServerConfig struct {
	Host string `envconfig:"HOST"`
}

type QueueConfig struct {
	Workers     int           `envconfig:"WORKERS"`
	QueueSize   int           `envconfig:"QUEUE_SIZE"`
	BaseBackoff time.Duration `envconfig:"BASE_BACKOFF"`
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
		return "0.0.0.0:8080"
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

	//var cfg Config
	//
	//local := time.FixedZone("MSK", 3*60*60)
	//time.Local = local
	//
	//if err := envconfig.Process("", &cfg); err != nil {
	//	panic(err)
	//}
	//if cfg.ServerConfig.Host == "" {
	//	panic(errors.New("cfg is required"))
	//}
	//
	//return &cfg
}

func getEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}
