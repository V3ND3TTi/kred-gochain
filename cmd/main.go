package main

import (
	"fmt"
	"github.com/V3ND3TTi/kred-gochain/blockchain"
	"math/big"
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

	fmt.Println("\n⛏️  Simulating participation rewards...")
	reward := blockchain.Kred(10) // 10 Kred

	for i := 1; i <= 5; i++ {
		recipient := blockchain.GetNextParticipant()
		if recipient == nil {
			fmt.Println("❌ No active participant found.")
			continue
		}

		// Credit reward to recipient
		blockchain.AdjustBalance(recipient.Address, reward)

		tx := blockchain.Transaction{
			Sender:    "NETWORK",
			Recipient: recipient.Address,
			Amount:    new(big.Int).Set(reward),
		}

		chain.AddBlock([]blockchain.Transaction{tx}, reward)
		fmt.Printf("🎉 Block #%d → %s rewarded 10 Kred\n", i, recipient.Address)
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
