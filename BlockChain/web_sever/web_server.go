package websever

import (
	"bufio"
	"encoding/json"
	blockchain "go_learn/block_chain"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

var bcServer chan []*blockchain.Block

// Getenv is a function that initializes the environment variables and starts a server to handle incoming connections.
func Getenv(bc *blockchain.Blockchain) {
	// Create a channel to receive blocks from the blockchain.
	bcServer = make(chan []*blockchain.Block)

	// Load the environment variables from the web.env file.
	err := godotenv.Load("./web_sever/web.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Print the value of the ADDR environment variable.
	log.Printf("%s lives in", os.Getenv("ADDR"))

	// Start a TCP server on the specified address.
	server, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	// Accept incoming connections and handle them concurrently.
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Get the transaction from the connection and add it to the blockchain.
		getTransaction(conn, bc)

		// Handle the connection concurrently.
		go handleConn(conn, bc)
	}
}

// handleConn is a function that handles the connection with a client.
// It receives the connection object and a pointer to a blockchain object.
func handleConn(conn net.Conn, bc *blockchain.Blockchain) {
	// output stores the JSON representation of the blockchain blocks
	var output []byte

	// Start a new goroutine that runs in the background
	go func() {
		// Loop indefinitely with a sleep of 30 seconds
		for {
			time.Sleep(30 * time.Second)

			// Marshal the blockchain blocks into JSON format
			output, _ = json.Marshal(bc.Blocks)

			// Write the JSON output to the connection
			io.WriteString(conn, string(output))
		}
	}()

	// Loop over the bcServer channel
	for range bcServer {
		// Print the blockchain blocks using spew.Dump
		spew.Dump(bc.Blocks)
	}

	// Close the connection when the function ends
	defer conn.Close()
}

// Refactored function to get a transaction
func getTransaction(conn net.Conn, bc *blockchain.Blockchain) {
	// Prompt the client to enter a new data
	io.WriteString(conn, "Enter a new Data:")

	// Create a scanner to read the input from the client
	scanner := bufio.NewScanner(conn)

	// Start a new goroutine to handle the client input
	go func() {
		// Read the input from the client line by line
		for scanner.Scan() {
			// Get the data entered by the client
			data := scanner.Text()

			// Add the data to the blockchain
			bc.AddBlock(data)

			// Print the length of the blockchain
			log.Println(len(bc.Blocks))

			// Replace the blockchain with the updated blocks
			blockchain.ReplaceChain(bc, bc.Blocks)

			// Send the updated blocks to the blockchain server
			bcServer <- bc.Blocks

			// Prompt the client to enter a new data
			io.WriteString(conn, "\nEnter a new Data:")
		}
	}()
}
