package mining

const (
	// SubsidyHalvingInterval is the interval (in terms of blocks mined) for
	// which the block subsidy is halved.
	SubsidyHalvingInterval uint64 = 210000

	// InitialSubsidySatoshis is the initial block subsidy in satoshis.
	InitialSubsidySatoshis uint64 = 50 * 1000000000

	// TargetSize is the size in bytes of the block difficulty.
	TargetSize = 32
)

// CalcBlockSubsidy calculates the block subsidy given the height of a block.
func CalcBlockSubsidy(height uint64) uint64 {
	halvings := height / SubsidyHalvingInterval
	if halvings >= 64 {
		return 0
	}

	sub := InitialSubsidySatoshis
	return sub >> halvings
}
