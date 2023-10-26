package main

import (
	"fmt"
)

// main is the entry point of the program.
func main() {
	// Create the genesis block.
	genesisBlock := NewGenesisBlock()

	// Print the data of the genesis block.
	fmt.Printf("Data: %s\n", genesisBlock.Data)

	// Print the previous hash of the genesis block.
	fmt.Printf("Prev. Hash: %x\n", genesisBlock.PrevHash)

	// Print the hash of the genesis block.
	fmt.Printf("Hash: %x\n", genesisBlock.Hash)
}
