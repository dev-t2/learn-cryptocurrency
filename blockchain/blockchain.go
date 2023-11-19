package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Height   int    `json:"height"`
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
}

type blockchain struct {
	blocks []*Block
}

var ErrNotFound = errors.New("Block Not Found")

var b *blockchain

var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))

	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)

	if totalBlocks == 0 {
		return ""
	}

	return GetBlockchain().blocks[totalBlocks - 1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{len(GetBlockchain().blocks) + 1, data, "", getLastHash()}

	newBlock.calculateHash()

	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func () {
			b = &blockchain{}

			b.AddBlock("Genesis Block")
		})
	}

	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

func (b *blockchain) Block(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}

	return b.blocks[height - 1], nil
}