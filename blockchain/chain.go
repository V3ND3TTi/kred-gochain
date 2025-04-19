package blockchain

import (
	"math/big"
	"time"
)

// Blockchain contains the slice of all validated blocks
type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain initializes the chain with the genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{GenesisBlock()},
	}
}

// LatestBlock returns the most recent block
func (bc *Blockchain) LatestBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

// AddBlock creates a new block and adds it to the chain
func (bc *Blockchain) AddBlock(txs []Transaction) {
	prevBlock := bc.LatestBlock()
	reward := CalculateReward(prevBlock.Index + 1)

	newBlock := &Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now(),
		Transactions: txs,
		PrevHash:     prevBlock.Hash,
		Nonce:        0,
		Reward:       new(big.Int).Set(reward),
	}

	newBlock.MerkleRoot = CalculateMerkleRoot(newBlock.Transactions)
	newBlock.Hash = CalculateHash(*newBlock)

	bc.Blocks = append(bc.Blocks, newBlock)
}

// IsValid checks the integrity of the entire blockchain
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		if current.Hash != CalculateHash(*current) {
			return false
		}

		if current.PrevHash != previous.Hash {
			return false
		}
	}
	return true
}
