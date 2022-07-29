package destiny2

import (
	"fmt"
	"log"

	"github.com/crit/raid-helper/worker/internal/bungie"
)

type (
	Inventory struct {
		Category string
		Items    []Item
	}

	Item struct {
		Name string
		Tier string
		Type string
		Icon string
	}
)

func Items(client *bungie.Client, inventory *Inventory, salesItemCategories map[string]any) error {
	raw := map[string]any{}

	err := client.Get(&raw, fmt.Sprintf("/Manifest/%d/%s/", InventoryItemDefinitionType, "item hash string"))
	if err != nil {
		return err
	}

	for key := range raw {
		log.Println(key)
	}

	return nil
}
