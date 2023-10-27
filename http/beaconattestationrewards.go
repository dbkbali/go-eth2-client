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
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	api "github.com/dbkbali/go-eth2-client/api/v1"
	"github.com/dbkbali/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

type apiRewardsResponseJSON struct {
	Data *struct {
		TotalRewards []struct {
			ValidatorIndex phase0.ValidatorIndex `json:"validator_index"`
			Head           phase0.Gwei           `json:"head"`
			Target         phase0.Gwei           `json:"target"`
			Source         phase0.Gwei           `json:"source"`
			Inactivity     phase0.Gwei           `json:"inactivity"`
		} `json:"total_rewards"`
	} `json:"data"`
}

// BeaconAttestationRewards returns the attestation rewards for the given epoch.for
// the designated validators
// validatorIndices is a list of validators to restrict the returned values.  At least one validator must be provided for a result to be returned
func (s *Service) BeaconAttestationRewards(ctx context.Context, requestEpoch phase0.Epoch, validatorIndices []phase0.ValidatorIndex) (map[phase0.ValidatorIndex]*api.BeaconAttestationRewards, error) {
	specJSON, err := json.Marshal(validatorIndices)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal JSON")
	}

	url := fmt.Sprintf("/eth/v1/beacon/rewards/attestations/%d", requestEpoch)
	if len(validatorIndices) != 0 {
		ids := make([]string, len(validatorIndices))
		for i := range validatorIndices {
			ids[i] = fmt.Sprintf("%d", validatorIndices[i])
		}
		// url = fmt.Sprintf("%s?id=%s", url, strings.Join(ids, ","))
	}

	respBodyReader, err := s.post(ctx, url, bytes.NewBuffer(specJSON))
	if err != nil {
		return nil, errors.Wrap(err, "failed to request beacon attestation rewards")
	}
	if respBodyReader == nil {
		return nil, errors.New("failed to obtain beaconAttestorRewards")
	}
	var beaconAttestorRewardsJSON apiRewardsResponseJSON
	if err := json.NewDecoder(respBodyReader).Decode(&beaconAttestorRewardsJSON); err != nil {
		return nil, errors.Wrap(err, "failed to parse beacon attestation rewards")
	}
	if beaconAttestorRewardsJSON.Data == nil {
		return nil, errors.New("no attestation rewards for validators returned")
	}

	res := make(map[phase0.ValidatorIndex]*api.BeaconAttestationRewards)
	for _, totalReward := range beaconAttestorRewardsJSON.Data.TotalRewards {
		if res[totalReward.ValidatorIndex] == nil {
			res[totalReward.ValidatorIndex] = &api.BeaconAttestationRewards{}
		}
		res[totalReward.ValidatorIndex].Index = totalReward.ValidatorIndex
		res[totalReward.ValidatorIndex].Epoch = requestEpoch
		res[totalReward.ValidatorIndex].Head = totalReward.Head
		res[totalReward.ValidatorIndex].Target = totalReward.Target
		res[totalReward.ValidatorIndex].Source = totalReward.Source
		res[totalReward.ValidatorIndex].Inactivity = totalReward.Inactivity
		res[totalReward.ValidatorIndex].Total = totalReward.Head + totalReward.Target + totalReward.Source + totalReward.Inactivity
	}
	return res, nil
}
