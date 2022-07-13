package xur

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const locationURL = "https://paracausal.science/xur/current.json"

type (
	LocationResult struct {
		Found    bool
		Position string
	}
)

func Location() (*LocationResult, error) {
	req, err := http.NewRequest(http.MethodGet, locationURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", locationURL, res.Status)
	}

	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// The "null" return has a carriage return.
	if strings.TrimSpace(string(raw)) == "null" {
		return &LocationResult{Found: false}, nil
	}

	// TODO: marshal raw to a standard response object

	return &LocationResult{Found: true}, nil
}
