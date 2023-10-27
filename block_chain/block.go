package block_chain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block represents a block in a blockchain.
type Block struct {
	Index     int64  // The index of the block in the chain.
	Timestamp int64  // The timestamp when the block was created.
	Data      []byte // The data stored in the block.
	PrevHash  []byte // The hash of the previous block in the chain.
	Hash      []byte // The hash of the current block.
}

// getHash calculates the hash of a given block.
func getHash(b *Block) []byte {
	// Convert the index to a byte slice.
	index := []byte(strconv.FormatInt(b.Index, 10))

	// Convert the timestamp to a byte slice.
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// Concatenate the previous hash, data, and timestamp.
	headers := bytes.Join([][]byte{index, b.PrevHash, b.Data, timestamp}, []byte{})

	// Calculate the SHA256 hash of the concatenated headers.
	hash := sha256.Sum256(headers)

	return hash[:]
}

// NewBlock creates a new block with the given data and previous hash.
// It returns a pointer to the newly created block.
func NewBlock(index int64, data string, prevHash []byte) *Block {
	// Create a new block object with the current index+1, timestamp, data, previous hash,
	// and an empty hash field.
	block := &Block{
		Index:     index + 1,
		Timestamp: time.Now().Unix(),
		Data:      []byte(data),
		PrevHash:  prevHash,
		Hash:      []byte{},
	}

	// Calculate the hash value of the block using the GetHash method
	block.Hash = getHash(block)

	return block
}

// NewGenesisBlock creates a new genesis block.
// A genesis block is the first block in a blockchain and does not have any preceding block.
func NewGenesisBlock() *Block {
	// Create a new block with the data "Genesis Block" and an empty byte slice as the previous hash.
	return NewBlock(0, "Genesis Block", []byte{})
}

// IsBlockValid checks if a block is valid
func IsBlockValid(prveBlock *Block, newBlock *Block) bool {

	// Compare the index of the previous block to the index of the new block.
	if prveBlock.Index+1 != newBlock.Index {
		return false
	}

	// Compare the hash of the previous block to the hash of the new block.
	if !bytes.Equal(prveBlock.Hash, newBlock.PrevHash) {
		return false
	}

	// Compare the hash of the previous block to the hash of the previous block.
	if !bytes.Equal(getHash(prveBlock), prveBlock.Hash) {
		return false
	}

	// Compare the hash of the new block to the hash of the new block.
	if !bytes.Equal(getHash(newBlock), newBlock.Hash) {
		return false
	}

	return true
}
