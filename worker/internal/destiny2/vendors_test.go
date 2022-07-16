package destiny2_test

import (
	"testing"

	"github.com/crit/raid-helper/worker/internal/bungie"
	"github.com/crit/raid-helper/worker/internal/destiny2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVendors(t *testing.T) {
	client := bungie.NewClient(bungie.ConfigFromEnv())

	res, err := destiny2.Vendors(client)

	require.Nil(t, err)
	assert.Equal(t, `{"unknown":""}`, res)
}
