package blockchain

import (
	"crypto/sha256"
	"fmt"
)

// CalculateMerkleRoot returns the Merkle root of the provided transactions.
func CalculateMerkleRoot(txs []Transaction) string {
	if len(txs) == 0 {
		return ""
	}

	var hashes []string
	for _, tx := range txs {
		hashes = append(hashes, HashTransaction(tx))
	}

	for len(hashes) > 1 {
		var newLevel []string

		for i := 0; i < len(hashes); i += 2 {
			if i+1 < len(hashes) {
				combined := hashes[i] + hashes[i+1]
				newLevel = append(newLevel, hashString(combined))
			} else {
				// Duplicate last hash if odd number of items
				combined := hashes[i] + hashes[i]
				newLevel = append(newLevel, hashString(combined))
			}
		}

		hashes = newLevel
	}

	return hashes[0] // Final Merkle root
}

// hashString returns the SHA-256 hash of a string in hex format.
func hashString(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
