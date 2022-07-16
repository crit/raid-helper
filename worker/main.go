package main

import (
	"flag"
	"log"
	"time"

	"github.com/crit/raid-helper/worker/internal"
	"github.com/crit/raid-helper/worker/internal/storage"
	"github.com/crit/raid-helper/worker/internal/xur"
	"github.com/robfig/cron/v3"
)

var (
	store storage.FileStorage
)

// ./xur-worker xur-worker.env
func main() {
	var err error

	flag.Parse()

	// load env file
	err = internal.LoadEnvFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	// Setup Remote Storage
	store, err = storage.NewS3("raid-help.critrussell.com", "us-west-2")
	if err != nil {
		log.Fatal(err)
	}

	// Setup timezone for jobs
	loc, err := time.LoadLocation("MST")
	if err != nil {
		log.Fatal(err)
	}

	// Setup cron system
	c := cron.New(
		cron.WithLocation(loc),
		cron.WithLogger(cron.VerbosePrintfLogger(log.Default())),
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
	)

	// Run at 3am on Friday every week
	_, err = c.AddFunc("0 3 * * FRI", writeXurLocation)
	if err != nil {
		log.Fatal(err)
	}

	// Run at 3am on Friday every week
	// _, err = c.AddFunc("0 03 * * FRI", writeXurInventory)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Start process (blocking)
	c.Run()
}

func writeXurLocation() {
	res, err := xur.Location()
	if err != nil {
		log.Println(err.Error())
		return
	}

	if res == nil {
		return
	}

	err = store.Set("data/xur-location.json", res.Bytes())
	if err != nil {
		log.Printf("unable to write xur-location.json: %v", err)
	}
}

func writeXurInventory() {
	res, err := xur.Inventory()
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(res)

	// TODO: write json file to s3 with data for the frontend about xur's inventory
}
