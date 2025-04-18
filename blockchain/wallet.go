package blockchain

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// Wallet represents a simple address + balance container.
type Wallet struct {
	Address string
	Balance uint64 // In Koins
}

var wallets = make(map[string]*Wallet) // In-memory wallet storage

// CreateWallet generates a new wallet with a random address and optional starting balance.
func CreateWallet(startingBalance uint64) *Wallet {
	address := generateAddress()
	wallet := &Wallet{
		Address: address,
		Balance: startingBalance,
	}

	wallets[address] = wallet
	return wallet
}

// GetWallet returns a wallet by address (if it exists).
func GetWallet(address string) (*Wallet, bool) {
	wallet, exists := wallets[address]
	return wallet, exists
}

// AdjustBalance adds/subtracts balance from a wallet.
func AdjustBalance(address string, amount int64) bool {
	wallet, exists := wallets[address]
	if !exists {
		return false
	}

	// Prevent negative balance
	if amount < 0 && wallet.Balance < uint64(-amount) {
		return false
	}

	wallet.Balance = uint64(int64(wallet.Balance) + amount)
	return true
}

// ListWallets prints all wallets and their balances.
func ListWallets() {
	fmt.Println("📒 Wallets:")
	for _, w := range wallets {
		fmt.Printf("  → %s: %d Koins (%.4f Kred)\n", w.Address, w.Balance, float64(w.Balance)/1e18)
	}
}

// generateAddress creates a pseudo-random 20-byte address (like Ethereum-style)
func generateAddress() string {
	b := make([]byte, 20)
	_, _ = rand.Read(b)
	return "KRDx" + hex.EncodeToString(b)
}
