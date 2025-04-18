package blockchain

import (
	"crypto/sha256"
	"fmt"
)

// Transaction represents a transfer of Koins between two addresses.
type Transaction struct {
	Sender    string
	Recipient string
	Amount    uint64 // In Koins (smallest unit)
}

// HashTransaction returns the SHA-256 hash of a transaction's data.
func HashTransaction(tx Transaction) string {
	raw := fmt.Sprintf("%s:%s:%d", tx.Sender, tx.Recipient, tx.Amount)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(raw)))
}
