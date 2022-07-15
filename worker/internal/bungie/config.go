package bungie

import (
	"log"
	"net/http"

	"github.com/crit/raid-helper/worker/internal/bind"
	"github.com/crit/raid-helper/worker/internal/errors"
)

type (
	Config struct {
		ApiKey string `env:"BUNGIE_API_KEY"`
	}
)

func ConfigFromEnv() *Config {
	var cfg Config
	err := bind.Env(&cfg)
	if err != nil {
		e := errors.New(http.StatusInternalServerError, "error binding config: %v", err)
		log.Fatalf("%s", e.Error())
	}

	return &cfg
}
