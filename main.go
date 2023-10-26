package main

import (
	"fmt"
)

func main() {
	// 创建起源块
	genesisBlock := NewGenesisBlock()

	// 打印起源块的内容
	fmt.Printf("Data: %s\n", genesisBlock.Data)
	fmt.Printf("Prev. Hash: %x\n", genesisBlock.PrevHash)
	fmt.Printf("Hash: %x\n", genesisBlock.Hash)
}
