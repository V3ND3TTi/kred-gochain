package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []Transaction
	MerkleRoot   string
	PrevHash     string
	Hash         string
	Nonce        int
	Reward       uint64 // In Koins (smallest unit)
}

// NewBlock creates a new block and returns a pointer to it.
func NewBlock(index int, txs []Transaction, prevHash string, nonce int, reward uint64) *Block {
	timestamp := time.Now()
	merkleRoot := CalculateMerkleRoot(txs)

	block := &Block{
		Index:        index,
		Timestamp:    timestamp,
		Transactions: txs,
		MerkleRoot:   merkleRoot,
		PrevHash:     prevHash,
		Nonce:        nonce,
		Reward:       reward,
	}

	block.Hash = CalculateHash(*block)
	return block
}

// CalculateHash returns the SHA-256 hash of a block's core fields.
func CalculateHash(b Block) string {
	record := fmt.Sprintf("%d%s%s%s%d%d",
		b.Index,
		b.Timestamp.Format(time.RFC3339Nano),
		b.MerkleRoot,
		b.PrevHash,
		b.Nonce,
		b.Reward,
	)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(record)))
}

// GenesisBlock creates the first block in the chain.
func GenesisBlock() *Block {
	genesisTx := Transaction{
		Sender:    "GENESIS",
		Recipient: "KoinLab",
		Amount:    10_000_000_000_000_000_000, // 10 Kred
	}

	txs := []Transaction{genesisTx}
	return NewBlock(0, txs, "0", 0, genesisTx.Amount)
}
