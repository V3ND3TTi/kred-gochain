package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Nonce        int
	MerkleRoot   string
	Reward       *big.Int
}

// CalculateHash generates a SHA256 hash of the block
func CalculateHash(b Block) string {
	txData := ""
	for _, tx := range b.Transactions {
		txData += tx.Sender + tx.Recipient + tx.Amount.String()
	}

	blockData := fmt.Sprintf("%d%s%s%s%d%s%s",
		b.Index,
		b.Timestamp.String(),
		txData,
		b.PrevHash,
		b.Nonce,
		b.MerkleRoot,
		b.Reward.String(),
	)

	hash := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hash[:])
}

// GenesisBlock creates the first block in the chain
func GenesisBlock() *Block {
	reward := Kred(10) // 10 Kred = 1e18 * 10 Koins

	genesisTx := Transaction{
		Sender:    "NETWORK",
		Recipient: "KRDxGENESIS",
		Amount:    new(big.Int).Set(reward),
	}

	block := &Block{
		Index:        0,
		Timestamp:    time.Now(),
		Transactions: []Transaction{genesisTx},
		PrevHash:     "0",
		Hash:         "",
		Nonce:        0,
		MerkleRoot:   "",
		Reward:       new(big.Int).Set(reward),
	}

	block.MerkleRoot = CalculateMerkleRoot(block.Transactions)
	block.Hash = CalculateHash(*block)

	return block
}

// PrintBlock displays the block data for CLI/debug
func PrintBlock(b *Block) {
	fmt.Printf("Block #: %d\n", b.Index)
	fmt.Printf("Timestamp: %s\n", b.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("PrevHash: %s\n", b.PrevHash)
	fmt.Printf("Hash: %s\n", b.Hash)
	fmt.Printf("Merkle Root: %s\n", b.MerkleRoot)
	fmt.Printf("Reward: %s Koins\n", b.Reward.String())
	fmt.Println("Transactions:")
	for _, tx := range b.Transactions {
		fmt.Printf("  → %s sent %s Koins to %s\n", tx.Sender, tx.Amount.String(), tx.Recipient)
	}
}

// IsHashValid checks if a hash meets the difficulty criteria
func IsHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}
