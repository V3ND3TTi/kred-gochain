package blockchain

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Wallet represents a basic address with a big integer balance
type Wallet struct {
	Address string
	Balance *big.Int
}

var wallets = make(map[string]*Wallet)
var activeNodes []*Wallet
var rewardIndex int

// CreateWallet generates a new wallet with 0 Kred
func CreateWallet() *Wallet {
	address := generateAddress()
	wallet := &Wallet{
		Address: address,
		Balance: big.NewInt(0),
	}

	wallets[address] = wallet
	return wallet
}

// GetWallet retrieves a wallet by address
func GetWallet(address string) (*Wallet, bool) {
	wallet, exists := wallets[address]
	return wallet, exists
}

// AdjustBalance adds a big.Int amount to the wallet balance
func AdjustBalance(address string, amount *big.Int) bool {
	wallet, exists := wallets[address]
	if !exists {
		return false
	}

	// Prevent negative balance (only if subtracting)
	if amount.Sign() < 0 {
		newBal := new(big.Int).Add(wallet.Balance, amount)
		if newBal.Sign() < 0 {
			return false // insufficient funds
		}
	}

	wallet.Balance.Add(wallet.Balance, amount)
	return true
}

// ListWallets prints all wallets and their balances
func ListWallets() {
	fmt.Println("💵 Wallets:")
	for _, w := range wallets {
		kred := new(big.Float).Quo(new(big.Float).SetInt(w.Balance), big.NewFloat(1e18))
		fmt.Printf("  → %s: %s Koins (%.4f Kred)\n", w.Address, w.Balance.String(), kred)
	}
}

// RegisterNode adds a wallet to the participation reward cycle
func RegisterNode(wallet *Wallet) {
	activeNodes = append(activeNodes, wallet)
}

// GetNextParticipant returns the next wallet in the reward cycle
func GetNextParticipant() *Wallet {
	if len(activeNodes) == 0 {
		return nil
	}
	w := activeNodes[rewardIndex%len(activeNodes)]
	rewardIndex++
	return w
}

// GetAllWallets returns the full wallet map
func GetAllWallets() map[string]*Wallet {
	return wallets
}

// generateAddress creates a pseudo-random 20-byte hex address
func generateAddress() string {
	b := make([]byte, 20)
	_, _ = rand.Read(b)
	return "KRDx" + hex.EncodeToString(b)
}

// Kred returns a big.Int of amount * 1e18 (Koins)
func Kred(amount int64) *big.Int {
	base := big.NewInt(amount)
	factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	return new(big.Int).Mul(base, factor)
}

// Transfer handles a wallet-to-wallet transaction
func Transfer(from, to string, amount *big.Int) bool {
	fromWallet, existsFrom := wallets[from]
	toWallet, existsTo := wallets[to]

	if !existsFrom || !existsTo {
		return false
	}

	// Ensure sender has enough balance
	if fromWallet.Balance.Cmp(amount) < 0 {
		return false
	}

	fromWallet.Balance.Sub(fromWallet.Balance, amount)
	toWallet.Balance.Add(toWallet.Balance, amount)

	return true
}
