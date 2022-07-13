package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type (
	FormattedResponse struct {
		URL  string `json:"url"`
		Data any    `json:"data"`
	}

	GetVendor struct {
		MembershipType string
		MembershipID   string
		CharacterID    string
		VendorHash     string
	}
)

func main() {
	url := "https://www.bungie.net/Platform/Destiny2/Vendors/?components=402"

	apiKey := os.Getenv("BUNGIE_API_KEY")

	if apiKey == "" {
		log.Fatal("missing BUNGIE_API_KEY env var")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-API-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("%s: %s", url, res.Status)
	}

	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data any

	err = json.Unmarshal(raw, &data)
	if err != nil {
		log.Fatal(err)
	}

	var fileData FormattedResponse

	fileData.URL = url
	fileData.Data = data

	formatted, err := json.MarshalIndent(&fileData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("responses/%s.json", time.Now().Format("2006-01-02-150405"))
	err = ioutil.WriteFile(name, formatted, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
