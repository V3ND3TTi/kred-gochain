package main

import (
	"fmt"

	"github.com/V3ND3TTi/kred-gochain/blockchain"
	"github.com/V3ND3TTi/kred-gochain/wallet"
)

func main() {
	// 🔗 Create chain instance
	chain := blockchain.NewBlockchain()

	// 👤 Create wallets
	alice := wallet.CreateWallet()
	bob := wallet.CreateWallet()
	charlie := wallet.CreateWallet()

	fmt.Println("🔐 Wallets:")
	fmt.Println("  Alice:", alice.Address)
	fmt.Println("  Bob:  ", bob.Address)
	fmt.Println("  Charlie:", charlie.Address)

	// 🏆 Simulate participation rewards
	fmt.Println("\n⛏️ Simulating 5 reward blocks...")
	for i := 0; i < 5; i++ {
		reward := blockchain.GetCurrentReward(len(chain.Blocks))
		for _, addr := range []string{alice.Address, bob.Address, charlie.Address} {
			wallet.AdjustBalance(addr, reward)

			// Add block per reward event
			tx := blockchain.Transaction{
				Sender:    "KRED_SYSTEM",
				Recipient: addr,
				Amount:    reward,
			}
			chain.AddBlock([]blockchain.Transaction{tx})
		}
	}

	// 💸 Simulate a transaction
	fmt.Printf("\n📤 Alice sends 2 Kred to Bob\n")
	amount := wallet.Kred(2)
	wallet.Transfer(alice.Address, bob.Address, amount)

	tx := blockchain.Transaction{
		Sender:    alice.Address,
		Recipient: bob.Address,
		Amount:    amount,
	}
	chain.AddBlock([]blockchain.Transaction{tx})

	// 🧾 Final balances
	fmt.Println("\n💵 Final Balances:")
	wallet.ListWallets()
}
