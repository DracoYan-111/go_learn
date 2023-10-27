package block_chain

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prveBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prveBlock.Index, data, prveBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// func GetBlockchain() *Blockchain {
// 	return blockchain
// }
