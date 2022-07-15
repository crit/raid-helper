package xur

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/crit/raid-helper/worker/internal/errors"
)

/*
{
  "location": "tower",
  "locationName": "The Tower Hangar",
  "placeHash": 3747705955,
  "destinationHash": 1272484305,
  "bubbleIndex": 3,
  "placeName": "Earth",
  "destinationName": "The Last City",
  "bubbleName": "Hangar"
}
*/
const locationURL = "https://paracausal.science/xur/current.json"

type (
	locationPayload struct {
		Location        string `json:"location"`
		LocationName    string `json:"locationName"`
		PlaceHash       int64  `json:"placeHash"`
		DestinationHash int64  `json:"destinationHash"`
		BubbleIndex     int    `json:"bubbleIndex"`
		PlaceName       string `json:"placeName"`
		DestinationName string `json:"destinationName"`
		BubbleName      string `json:"bubbleName"`
	}

	LocationResult struct {
		Expires         string `json:"expires"`
		DestinationName string `json:"destinationName"`
		LocationName    string `json:"locationName"`
	}
)

func Location() (*LocationResult, error) {
	req, err := http.NewRequest(http.MethodGet, locationURL, nil)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", locationURL, res.Status)
	}

	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	// The "null" return has a carriage return.
	if strings.TrimSpace(string(raw)) == "null" {
		return nil, nil
	}

	var payload locationPayload
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var result LocationResult

	result.Expires = time.Now().AddDate(0, 0, 3).Format("2006-01-02")
	result.DestinationName = payload.DestinationName
	result.LocationName = payload.LocationName

	return &result, nil
}

func (res LocationResult) Bytes() []byte {
	raw, _ := json.Marshal(res)
	return raw
}
