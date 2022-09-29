package destiny2

import (
	"log"

	"github.com/crit/raid-helper/worker/internal/bungie"
)

type (
	User struct {
		ID       string
		Username string
		Platform string
	}
)

func UserData(client *bungie.Client, username, platform string) (*User, error) {
	var res bungie.Result

	err := client.Get(&res, searchDestinyPlayerURL(username, platform))
	if err != nil {
		return nil, err
	}

	log.Printf("%#v", res)

	return &User{
		ID:       "id",
		Username: username,
		Platform: platform,
	}, nil
}
