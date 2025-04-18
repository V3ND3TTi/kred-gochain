package blockchain

// Blockchain represents the full Kred chain.
type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain initializes a new chain with the Genesis block.
func NewBlockchain() *Blockchain {
	genesis := GenesisBlock()
	return &Blockchain{
		Blocks: []*Block{genesis},
	}
}

// AddBlock creates and appends a new block to the chain.
func (bc *Blockchain) AddBlock(txs []Transaction, reward uint64) *Block {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(lastBlock.Index+1, txs, lastBlock.Hash, 0, reward)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

// LatestBlock returns the most recent block in the chain.
func (bc *Blockchain) LatestBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

// IsValid checks the integrity of the blockchain.
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		prev := bc.Blocks[i-1]
		curr := bc.Blocks[i]

		// Validate hash chain
		if curr.PrevHash != prev.Hash {
			return false
		}

		// Validate current block hash
		if curr.Hash != CalculateHash(*curr) {
			return false
		}

		// Validate Merkle root
		if curr.MerkleRoot != CalculateMerkleRoot(curr.Transactions) {
			return false
		}
	}

	return true
}
