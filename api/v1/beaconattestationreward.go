package v1

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

type BeaconAttestationRewards struct {
	Index      phase0.ValidatorIndex
	Pubkey     phase0.BLSPubKey `ssz-size:"48"`
	Epoch      phase0.Epoch
	Head       phase0.Gwei
	Target     phase0.Gwei
	Source     phase0.Gwei
	Inactivity phase0.Gwei
	Total      phase0.Gwei
}

type BeaconAttestationRewardsJSON struct {
	Index      string `json:"index"`
	Pubkey     string `json:"pubkey"`
	Epoch      string `json:"epoch"`
	Head       string `json:"head"`
	Target     string `json:"target"`
	Source     string `json:"source"`
	Inactivity string `json:"inactivity"`
	Total      string `json:"total"`
}

func (b *BeaconAttestationRewards) MarshalJSON() ([]byte, error) {
	return json.Marshal(&BeaconAttestationRewardsJSON{
		Index:      fmt.Sprintf("%d", b.Index),
		Pubkey:     fmt.Sprintf("%#x", b.Pubkey),
		Epoch:      fmt.Sprintf("%d", b.Epoch),
		Head:       fmt.Sprintf("%d", b.Head),
		Target:     fmt.Sprintf("%d", b.Target),
		Source:     fmt.Sprintf("%d", b.Source),
		Inactivity: fmt.Sprintf("%d", b.Inactivity),
		Total:      fmt.Sprintf("%d", b.Total),
	})
}

func (b *BeaconAttestationRewards) UnmarshalJSON(input []byte) error {
	var data BeaconAttestationRewardsJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}
	return b.unpack(&data)
}

func (b *BeaconAttestationRewards) unpack(data *BeaconAttestationRewardsJSON) error {

	if data.Index == "" {
		return errors.New("index missing")
	}
	index, error := strconv.ParseUint(data.Index, 10, 64)
	if error != nil {
		return errors.Wrap(error, "invalid value for index")
	}
	b.Index = phase0.ValidatorIndex(index)

	if data.Pubkey == "" {
		return errors.New("public key missing")
	}
	pubKey, err := hex.DecodeString(strings.TrimPrefix(data.Pubkey, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid value for public key")
	}
	if len(pubKey) != publicKeyLength {
		return errors.New("incorrect length for public key")
	}
	copy(b.Pubkey[:], pubKey)

	if data.Epoch == "" {
		return errors.New("request epoch missing")
	}
	epoch, err := strconv.ParseUint(data.Epoch, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for epoch")
	}
	b.Epoch = phase0.Epoch(epoch)

	head, err := strconv.ParseUint(data.Head, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for head reward")
	}
	b.Head = phase0.Gwei(head)

	target, err := strconv.ParseUint(data.Target, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for target reward")
	}
	b.Target = phase0.Gwei(target)

	source, err := strconv.ParseUint(data.Source, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for source reward")
	}
	b.Source = phase0.Gwei(source)

	inactivity, err := strconv.ParseUint(data.Inactivity, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for inactivity reward")
	}
	b.Inactivity = phase0.Gwei(inactivity)

	total, err := strconv.ParseUint(data.Total, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for total reward")
	}
	b.Total = phase0.Gwei(total)

	return nil
}

// String returns a string version of the structure.
func (b *BeaconAttestationRewards) String() string {
	data, err := json.Marshal(b)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}
	return string(data)
}
