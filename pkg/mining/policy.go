package mining

const (
	// RewardHalvingInterval is the interval (in terms of blocks mined) for
	// which the block reward is halved.
	RewardHalvingInterval uint64 = 210000

	// InitialRewardSatoshis is the initial block reward in satoshis.  It is
	// used as a base value to calculate block rewards.
	InitialRewardSatoshis uint64 = 5000000000
)

// BlockReward calculates the block reward given the height of a block.
func BlockReward(height uint64) uint64 {
	halvings := height / RewardHalvingInterval
	if halvings >= 64 {
		return 0
	}

	reward := InitialRewardSatoshis
	reward >>= halvings
	return reward
}
