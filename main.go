package main

import (
	"fmt"
	blockchain "go_learn/block_chain"
)

// main is the entry point of the program.
func main() {
	// Create a new instance of the Blockchain struct
	bc := blockchain.NewBlockchain()

	// Add a new block to the blockchain with the data "Send 1 BTC to Ivan"
	bc.AddBlock("Send 1 BTC to Ivan")

	// Add another block to the blockchain with the data "Send 2 more BTC to Ivan"
	bc.AddBlock("Send 2 more BTC to Ivan")

	// Iterate over each block in the blockchain
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)

		// Print the previous hash of the block
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)

		// Print the data stored in the block
		fmt.Printf("Data: %s\n", block.Data)

		// Print the hash of the block
		fmt.Printf("Hash: %x\n", block.Hash)

		// Print the timestamp of the block
		fmt.Printf("Timestamp: %d\n", block.Timestamp)

		// Print a new line to separate each block
		fmt.Println()
	}
}
