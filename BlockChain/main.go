package main

import (
	//"fmt"
	blockchain "go_learn/block_chain"
	websever "go_learn/web_sever"
	"log"
)

// main is the entry point of the program.
func main() {

	// Create a new instance of the Blockchain struct
	bc := blockchain.NewBlockchain()
	log.Println(len(bc.Blocks))

	websever.Getenv(bc)
}
