package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block represents a block in a blockchain.
type Block struct {
	Timestamp int64  // The timestamp when the block was created.
	Data      []byte // The data stored in the block.
	PrevHash  []byte // The hash of the previous block in the chain.
	Hash      []byte // The hash of the current block.
}

// SetHash calculates and sets the hash value of the block.
func (b *Block) SetHash() {
	// Convert the timestamp to a byte slice.
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// Concatenate the previous hash, data, and timestamp.
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp}, []byte{})

	// Calculate the hash using SHA256.
	hash := sha256.Sum256(headers)

	// Set the hash value of the block.
	b.Hash = hash[:]
}

// NewBlock creates a new block with the given data and previous hash.
// It returns a pointer to the newly created block.
func NewBlock(data string, prevHash []byte) *Block {
	// Create a new block object with the current timestamp, data, previous hash,
	// and an empty hash field.
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data:      []byte(data),
		PrevHash:  prevHash,
		Hash:      []byte{},
	}

	// Calculate and set the hash of the block using the SetHash method.
	block.SetHash()

	return block
}

// NewGenesisBlock creates a new genesis block.
// A genesis block is the first block in a blockchain and does not have any preceding block.
func NewGenesisBlock() *Block {
	// Create a new block with the data "Genesis Block" and an empty byte slice as the previous hash.
	return NewBlock("Genesis Block", []byte{})
}
