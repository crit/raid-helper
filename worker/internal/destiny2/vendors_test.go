package destiny2_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/crit/raid-helper/worker/internal/bungie"
	"github.com/crit/raid-helper/worker/internal/destiny2"
	"github.com/stretchr/testify/require"
)

func TestVendors(t *testing.T) {
	client := bungie.NewClient(bungie.ConfigFromEnv())

	res, err := destiny2.Vendors(client)
	require.Nil(t, err)

	raw, err := json.Marshal(res)
	require.Nil(t, err)

	err = ioutil.WriteFile("vendor-response.json", raw, 0644)
	require.Nil(t, err)
}
