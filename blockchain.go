package main

import (
	"bytes"
	"crypto/sha256"
)

type Blockchain struct {
	blocks []Block
}

type Block struct {
	previousBlockHash []byte
	transaction       Transaction
}

func (c *Blockchain) AddBlock(b *Block) *Block {
	// TODO: Add block.
	return nil
}

func (c *Blockchain) GetGenesisBlock() []byte {
	if len(c.blocks) == 0 {
		return nil
	}

	return c.blocks[0].getBlockHash()
}

func (b *Block) getBlockHash() []byte {
	buff := new(bytes.Buffer)

	if b.previousBlockHash != nil {
		buff.Write(b.previousBlockHash)
	}

	buff.Write(b.transaction.Hash())

	hash := sha256.Sum256(buff.Bytes())

	return hash[:]
}