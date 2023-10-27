package http_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	client "github.com/dbkbali/go-eth2-client"
	"github.com/dbkbali/go-eth2-client/http"
	"github.com/dbkbali/go-eth2-client/spec/phase0"
	"github.com/stretchr/testify/require"
)

func TestBeaconAttestationRewards(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		name              string
		epoch             phase0.Epoch
		expectedErrorCode int
		validatorIndices  []phase0.ValidatorIndex
	}{
		{
			name:              "Good",
			epoch:             238572,
			expectedErrorCode: 0,
			validatorIndices:  []phase0.ValidatorIndex{512735},
		},
	}

	service, err := http.New(ctx,
		http.WithTimeout(timeout),
		http.WithAddress(os.Getenv("HTTP_ADDRESS")),
	)
	require.NoError(t, err)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beaconAttestorRewards, err := service.(client.BeaconAttestationRewardsProvider).BeaconAttestationRewards(ctx, test.epoch, test.validatorIndices)
			if test.expectedErrorCode != 0 {
				require.Contains(t, err.Error(), fmt.Sprintf("%d", test.expectedErrorCode))
			} else {
				require.NoError(t, err)
				require.NotNil(t, beaconAttestorRewards)
			}
		})
	}
}
