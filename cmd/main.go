package main

import (
	"fmt"
	"github.com/V3ND3TTi/kred-gochain/blockchain"
	"time"
)

func main() {
	fmt.Println("🧱 Initializing node wallets...")

	// Create wallets
	w1 := blockchain.CreateWallet()
	w2 := blockchain.CreateWallet()
	w3 := blockchain.CreateWallet()

	fmt.Printf("W1: %s\n", w1.Address)
	fmt.Printf("W2: %s\n", w2.Address)
	fmt.Printf("W3: %s\n", w3.Address)

	// Register for participation
	blockchain.RegisterNode(w1)
	blockchain.RegisterNode(w2)
	blockchain.RegisterNode(w3)

	fmt.Println("\n✅ Participation nodes registered.")
	blockchain.ListWallets()

	// Initialize blockchain
	chain := blockchain.NewBlockchain()

	// Simulate active participants
	var activeNodes []*blockchain.Node
	for addr := range blockchain.GetAllWallets() {
		activeNodes = append(activeNodes, &blockchain.Node{
			Address:       addr,
			LastCheckIn:   time.Now(), // Pretend they just checked in
			Participating: true,
		})
	}

	// Simulate blocks with reward splits
	fmt.Println("\n⛏️ Simulating reward distribution for 5 blocks...")
	for i := 1; i <= 5; i++ {
		blockHeight := len(chain.Blocks)
		reward := blockchain.CalculateReward(blockHeight)

		rewards := blockchain.DistributeRewardEvenly(reward, activeNodes)

		// Apply rewards
		for addr, amount := range rewards {
			blockchain.AdjustBalance(addr, amount)
		}

		// Create a dummy tx for block (optional, or build from rewards)
		txs := []blockchain.Transaction{}
		chain.AddBlock(txs)

		fmt.Printf("✅ Block #%d — Distributed %s Koins to %d nodes\n", blockHeight, reward.String(), len(activeNodes))
	}

	fmt.Println("\n💸 Simulating wallet-to-wallet transfer:")
	amount := blockchain.Kred(5) // 5 Kred

	success := blockchain.Transfer(w1.Address, w2.Address, amount)
	if success {
		fmt.Printf("✅ Transferred 5 Kred from %s to %s\n", w1.Address, w2.Address)
	} else {
		fmt.Println("❌ Transfer failed.")
	}

	fmt.Println("\n📒 Final Wallet Balances:")
	blockchain.ListWallets()

	fmt.Println("\n🧱 Latest Block:")
	blockchain.PrintBlock(chain.LatestBlock())

	fmt.Printf("\n🔎 Chain Valid? %v\n", chain.IsValid())
}
