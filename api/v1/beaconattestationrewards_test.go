package v1_test

import (
	"encoding/json"
	"testing"

	api "github.com/dbkbali/go-eth2-client/api/v1"
	"github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

func TestBeaconAttestationRewardsJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
		err      string
	}{
		{
			name:  "valid input",
			input: []byte(`{"index":"512735","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"234512","head":"2674","target":"4976","source":"2678","inactivity":"0","total":"10000"}`),
			err:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res api.BeaconAttestationRewards
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				assert.Equal(t, string(test.input), string(rt))
				assert.Equal(t, string(rt), res.String())
			}
		})
	}
}
