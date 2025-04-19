package blockchain

import (
	"math/big"
	"time"
)

const (
	StartingReward  = 10          // in Kred
	Decimals        = 18          // decimal precision
	HalvingInterval = 126_144_000 // blocks per halving (4 years)
	MaxHalvings     = 100         // theoretical max before reward = 0
)

type Node struct {
	Address       string
	LastCheckIn   time.Time
	Participating bool
}

// GetCurrentReward calculates the reward for the current block height based on halving schedule
func GetCurrentReward(blockHeight int) *big.Int {
	halfCount := blockHeight / HalvingInterval

	if halfCount >= MaxHalvings {
		return big.NewInt(0) // reward is done
	}

	exponent := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(halfCount)), nil)
	baseReward := big.NewInt(StartingReward)
	baseReward.Mul(baseReward, new(big.Int).Exp(big.NewInt(10), big.NewInt(Decimals), nil)) // convert to Koins

	return new(big.Int).Div(baseReward, exponent)
}

// CalculateReward returns the reward for a given block
func CalculateReward(height int) *big.Int {
	return GetCurrentReward(height)
}

// DistributeRewardEvenly splits the reward among all currently participating nodes
func DistributeRewardEvenly(totalReward *big.Int, participatingNodes []*Node) map[string]*big.Int {
	rewards := make(map[string]*big.Int)
	count := len(participatingNodes)
	if count == 0 {
		return rewards
	}

	split := new(big.Int).Div(totalReward, big.NewInt(int64(count)))

	for _, node := range participatingNodes {
		rewards[node.Address] = new(big.Int).Set(split)
	}

	return rewards
}
