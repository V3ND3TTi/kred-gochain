package main

import (
	"fmt"
	"github.com/V3ND3TTi/kred-gochain/blockchain"
)

func main() {
	chain := blockchain.NewBlockchain()

	fmt.Println("🧱 Genesis Block:")
	printBlock(chain.Blocks[0])

	txs := []blockchain.Transaction{
		{Sender: "KRDxAlice", Recipient: "KRDxBob", Amount: 1_000_000_000_000_000}, // 0.001 Kred
		{Sender: "KRDxBob", Recipient: "KRDxCharlie", Amount: 500_000_000_000_000}, // 0.0005 Kred
	}

	reward := uint64(10 * 1e18) // 10 Kred in Koins
	chain.AddBlock(txs, reward)

	fmt.Println("\n🧱 Latest Block:")
	printBlock(chain.LatestBlock())

	fmt.Printf("\n✅ Chain Valid? %v\n", chain.IsValid())

	fmt.Println("\n💳 Creating wallets...")

	alice := blockchain.CreateWallet(5 * 1e18) // 5 Kred
	bob := blockchain.CreateWallet(0)
	// charlie := blockchain.CreateWallet(0)

	fmt.Println("🎉 Wallets created:")
	blockchain.ListWallets()

	fmt.Printf("\n💸 Sending 1.5 Kred from %s to %s...\n", alice.Address, bob.Address)
	success := blockchain.AdjustBalance(alice.Address, -int64(1_500_000_000_000_000_000)) && // -1.5 Kred
		blockchain.AdjustBalance(bob.Address, 1_500_000_000_000_000_000) // +1.5 Kred

	if success {
		fmt.Println("✅ Transfer successful!")
	} else {
		fmt.Println("❌ Transfer failed!")
	}

	fmt.Println("\n📒 Updated Wallets:")
	blockchain.ListWallets()
}

// Helper function to print block details nicely
func printBlock(b *blockchain.Block) {
	fmt.Printf("Block #: %d\n", b.Index)
	fmt.Printf("Timestamp: %s\n", b.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("PrevHash: %s\n", b.PrevHash)
	fmt.Printf("Hash: %s\n", b.Hash)
	fmt.Printf("Merkle Root: %s\n", b.MerkleRoot)
	fmt.Printf("Reward: %d Koins (%.4f Kred)\n", b.Reward, float64(b.Reward)/1e18)
	fmt.Println("Transactions:")
	for _, tx := range b.Transactions {
		fmt.Printf("  → %s sent %d Koins (%.4f Kred) to %s\n",
			tx.Sender, tx.Amount, float64(tx.Amount)/1e18, tx.Recipient)
	}
}
