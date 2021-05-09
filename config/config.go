package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

const namespace = "hello-go"

// Specification struct to load environment variables
type Specification struct {
	LogLevel string `split_words:"true"`
	Memcache struct {
		TTL       time.Duration `required:"true"`
		Size      int64         `required:"true"`
		Prunesize uint32        `required:"true"`
	}
	Port uint32 `required:"true"`
}

var Config Specification

func init() {
	err := envconfig.Process(namespace, &Config)
	if err != nil {
		log.Fatal().Err(err)
	}
}
