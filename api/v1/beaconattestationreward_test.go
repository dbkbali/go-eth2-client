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
			input: []byte(`{"index":"1","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"100","target":"200","source":"300","inactivity":"400","total":"1000"}`),
			err:   "",
		},
		{
			name:  "missing index",
			input: []byte(`{"pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"100","target":"200","source":"300","inactivity":"400","total":"1000"}`),
			err:   "index missing",
		},
		{
			name:  "invalid index",
			input: []byte(`{"index":"invalid","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"100","target":"200","source":"300","inactivity":"400","total":"1000"}`),
			err:   "invalid value for index: strconv.ParseUint: parsing \"invalid\": invalid syntax",
		},
		{
			name:  "missing public key",
			input: []byte(`{"index":"2","epoch":"2","head":"100","target":"200","source":"300","inactivity":"400","total":"1000"}`),
			err:   "public key missing",
		},
		{
			name:  "invalid public key",
			input: []byte(`{"index":"1","pubkey":"0x000Z02030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"100","target":"200","source":"300","inactivity":"400","total":"1000"}`),
			err:   "invalid value for public key: encoding/hex: invalid byte: U+005A 'Z'",
		},
		{
			name:  "incorrect length for public key",
			input: []byte(`{"index":"1","pubkey":"0x00010203040506070809020a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"100","target":"200","source":"300","inactivity":"400","total":"1000"}`),
			err:   "incorrect length for public key",
		},
		{
			name:  "missing epoch",
			input: []byte(`{"index":"1","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","head":"100","target":"200","source":"300","inactivity":"400","total":"1000"}`),
			err:   "request epoch missing",
		},
		{
			name:  "invalid epoch",
			input: []byte(`{"index":"1","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"invalid","head":"100","target":"200","source":"300","inactivity":"400","total":"1000"}`),
			err:   "invalid value for epoch: strconv.ParseUint: parsing \"invalid\": invalid syntax",
		},
		{
			name:  "invalid head",
			input: []byte(`{"index":"1","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"invalid","target":"200","source":"300","inactivity":"400","total":"1000"}`),
			err:   "invalid value for head reward: strconv.ParseUint: parsing \"invalid\": invalid syntax",
		},
		{
			name:  "invalid target reward",
			input: []byte(`{"index":"1","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"100","target":"invalid","source":"300","inactivity":"400","total":"1000"}`),
			err:   "invalid value for target reward: strconv.ParseUint: parsing \"invalid\": invalid syntax",
		},
		{
			name:  "invalid source reward",
			input: []byte(`{"index":"1","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"100","target":"200","source":"invalid","inactivity":"400","total":"1000"}`),
			err:   "invalid value for source reward: strconv.ParseUint: parsing \"invalid\": invalid syntax",
		},
		{
			name:  "invalid inactivity reward",
			input: []byte(`{"index":"1","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"100","target":"200","source":"4000","inactivity":"invalid","total":"1000"}`),
			err:   "invalid value for inactivity reward: strconv.ParseUint: parsing \"invalid\": invalid syntax",
		},
		{
			name:  "invalid total reward",
			input: []byte(`{"index":"1","pubkey":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f","epoch":"2","head":"100","target":"200","source":"4000","inactivity":"400","total":"invalid"}`),
			err:   "invalid value for total reward: strconv.ParseUint: parsing \"invalid\": invalid syntax",
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
