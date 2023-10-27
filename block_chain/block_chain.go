package block_chain

type Blockchain struct {
	Blocks []*Block
}

// AddBlock adds a new block to the blockchain with the given data.
func (bc *Blockchain) AddBlock(data string) {
	// Get the previous block in the blockchain.
	previousBlock := bc.Blocks[len(bc.Blocks)-1]

	// Create a new block with the index, data, and hash of the previous block.
	newBlock := NewBlock(previousBlock.Index, data, previousBlock.Hash)

	// Check if the new block is valid by comparing it with the previous block.
	if IsBlockValid(previousBlock, newBlock) {
		// Append the new block to the blockchain.
		bc.Blocks = append(bc.Blocks, newBlock)
		replaceChain(bc, bc.Blocks)
	}
}

// NewBlockchain creates a new instance of the Blockchain struct.
func NewBlockchain() *Blockchain {
	// Create a new Blockchain instance and initialize it with a single block, the genesis block.
	blockchain := &Blockchain{
		Blocks: []*Block{NewGenesisBlock()},
	}
	// Return the newly created blockchain instance.
	return blockchain
}

// replaceChain replaces the blocks in the blockchain with a new chain if the new chain is longer.
func replaceChain(bc *Blockchain, chain []*Block) {
	// Check if the length of the new chain is greater than the length of the current blockchain's blocks.
	if len(chain) > len(bc.Blocks) {
		// If true, replace the blockchain's blocks with the new chain.
		bc.Blocks = chain
	}
}
