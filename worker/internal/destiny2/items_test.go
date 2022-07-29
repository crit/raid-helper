package destiny2

import (
	"testing"

	"github.com/crit/raid-helper/worker/internal/bungie"
	"github.com/stretchr/testify/require"
)

func TestItems(t *testing.T) {
	client := bungie.NewClient(bungie.ConfigFromEnv())
	inv := Inventory{}
	source := map[string]any{}

	err := Items(client, &inv, source)
	require.Nil(t, err)
}
