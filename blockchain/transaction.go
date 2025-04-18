package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

// Transaction represents a transfer of Koins from one wallet to another
type Transaction struct {
	Sender    string
	Recipient string
	Amount    *big.Int
}

// HashTransaction returns the SHA256 hash of the transaction data
func HashTransaction(tx Transaction) string {
	txData := tx.Sender + tx.Recipient + tx.Amount.String()
	hash := sha256.Sum256([]byte(txData))
	return hex.EncodeToString(hash[:])
}
