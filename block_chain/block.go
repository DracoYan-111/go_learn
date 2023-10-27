package block_chain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
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

// SetHash calculates and sets the hash value of the block.
func (b *Block) SetHash() {

	// Convert the timestamp to a byte slice.
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// Create a new byte slice to store the headers.
	index := make([]byte, 8)
	binary.LittleEndian.PutUint64(index, uint64(b.Index))
	// Concatenate the previous hash, data, and timestamp.
	headers := bytes.Join([][]byte{index, b.PrevHash, b.Data, timestamp}, []byte{})

	// Calculate the hash using SHA256.
	hash := sha256.Sum256(headers)

	// Set the hash value of the block.
	b.Hash = hash[:]
}

// NewBlock creates a new block with the given data and previous hash.
// It returns a pointer to the newly created block.
func NewBlock(index int64, data string, prevHash []byte) *Block {
	// Create a new block object with the current timestamp, data, previous hash,
	// and an empty hash field.
	block := &Block{
		Index:     index + 1,
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
	return NewBlock(0, "Genesis Block", []byte{})
}
