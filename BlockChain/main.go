package main

import (
	//"fmt"
	blockchain "block_chain/block_chain"
	websever "block_chain/web_sever"
	"log"
)

// main is the entry point of the program.
func main() {

	// Create a new instance of the Blockchain struct
	bc := blockchain.NewBlockchain()
	log.Println(len(bc.Blocks))

	websever.Getenv(bc)
}
