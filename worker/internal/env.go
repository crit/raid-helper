package internal

import (
	"net/http"

	"github.com/crit/raid-helper/worker/internal/errors"
	"github.com/subosito/gotenv"
)

func LoadEnvFile(path string) error {
	if path == "" {
		return errors.New(http.StatusBadRequest, "missing env file")
	}

	err := gotenv.Load(path)
	if err != nil {
		return errors.New(http.StatusBadRequest, "unable to load env file %s: %v", path, err)
	}

	return nil
}
