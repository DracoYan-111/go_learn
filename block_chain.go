package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prveBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prveBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// func GetBlockchain() *Blockchain {
// 	return blockchain
// }
