package bungie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/crit/raid-helper/worker/internal/errors"
)

type (
	Client struct {
		domain string
		apiURI string
		apiKey string
		http   *http.Client
	}
)

func NewClient(cfg *Config) *Client {
	if cfg == nil {
		log.Fatal(errors.New(http.StatusInternalServerError, "nil config passed").Error())
	}

	if cfg.ApiKey == "" {
		log.Fatal(errors.New(http.StatusInternalServerError, "config missing api key").Error())
	}

	return &Client{
		domain: "www.bungie.net",
		apiURI: "Platform",
		apiKey: cfg.ApiKey,
		http:   http.DefaultClient,
	}
}

func (c *Client) Get(receiver any, uri string) error {
	req, err := http.NewRequest(http.MethodGet, c.url(uri), nil)
	if err != nil {
		return errors.New(http.StatusInternalServerError, "%v", err)
	}

	c.addAuthHeader(req)

	res, err := c.http.Do(req)
	if err != nil {
		return errors.New(http.StatusInternalServerError, "%v", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		raw, _ := ioutil.ReadAll(res.Body)
		return errors.New(res.StatusCode, string(raw))
	}

	err = json.NewDecoder(res.Body).Decode(receiver)
	if err != nil {
		return errors.New(http.StatusBadRequest, "json decode error: %v", err)
	}

	return nil
}

func (c *Client) url(uri string) string {
	return "https://" + path.Join(c.domain, c.apiURI, strings.TrimLeft(uri, "/"))
}

func (c *Client) addAuthHeader(req *http.Request) {
	req.Header.Add("X-API-Key", c.apiKey)
}

func (c *Client) addJsonHeader(req *http.Request) {
	req.Header.Add("content-type", "application/json")
}
