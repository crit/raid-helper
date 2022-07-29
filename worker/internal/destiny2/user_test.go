package destiny2

import (
	"testing"

	"github.com/crit/raid-helper/worker/internal/bungie"
	"github.com/stretchr/testify/require"
)

func TestUserData(t *testing.T) {
	client := bungie.NewClient(bungie.ConfigFromEnv())

	user, err := UserData(client, "TemporalRisk#5669", "All")
	require.Nil(t, err)

	t.Log(user)
}
