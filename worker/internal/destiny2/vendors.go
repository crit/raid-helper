package destiny2

import "github.com/crit/raid-helper/worker/internal/bungie"

func Vendors(client *bungie.Client) (any, error) {
	var data any
	err := client.Get(&data, "/Destiny2/Vendors/?components=402")
	if err != nil {
		return nil, err
	}

	return data, nil
}
