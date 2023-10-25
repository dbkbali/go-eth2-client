// Copyright © 2020 - 2023 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	api "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

type beaconAttestorRewardsJSON struct {
	Data []*api.Validator `json:"data"`
}

// BeaconAttestationRewards returns the attestation rewards for the given epoch.for
// the designated validators
// validatorIndices is a list of validators to restrict the returned values.  At least one validator must be provided for a result to be returned
func (s *Service) BeaconAttestationRewards(ctx context.Context, requestEpoch phase0.Epoch, validatorIndices []phase0.ValidatorIndex) (map[phase0.ValidatorIndex]*api.Validator, error) {

	url := fmt.Sprintf("/eth/v1/beacon/rewards/attestations/%d", requestEpoch)
	if len(validatorIndices) != 0 {
		ids := make([]string, len(validatorIndices))
		for i := range validatorIndices {
			ids[i] = fmt.Sprintf("%d", validatorIndices[i])
		}
		url = fmt.Sprintf("%s?id=%s", url, strings.Join(ids, ","))
	}

	respBodyReader, err := s.get(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request validators")
	}
	if respBodyReader == nil {
		return nil, errors.New("failed to obtain validators")
	}

	var beaconAttestorRewardsJSON beaconAttestorRewardsJSON
	if err := json.NewDecoder(respBodyReader).Decode(&beaconAttestorRewardsJSON); err != nil {
		return nil, errors.Wrap(err, "failed to parse validators")
	}
	if beaconAttestorRewardsJSON.Data == nil {
		return nil, errors.New("no attestation rewards for validators returned")
	}

	res := make(map[phase0.ValidatorIndex]*api.Validator)
	for _, validator := range validatorsJSON.Data {
		res[validator.Index] = validator
	}
	return res, nil
}
