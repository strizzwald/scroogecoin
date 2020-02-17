package main

import (
	"bytes"
	"encoding/binary"
	"math/big"
)

type ConsumedCoin struct {
	coin Scroogecoin
	signature *Signature
}

func (c *ConsumedCoin) hash() []byte {
	buffer := new(bytes.Buffer)
	coinHash := c.coin.Hash()

	binary.Write(buffer, binary.BigEndian, c.signature.r)
	binary.Write(buffer, binary.BigEndian, c.signature.s)

	return append(buffer.Bytes(), coinHash...)
}

type Signature struct {
	r *big.Int
	s *big.Int
}

type Transaction interface {
	Hash() []byte
	ConsumedCoins() []ConsumedCoin
	CreatedCoins() []Scroogecoin
}

type CoinCreationTransaction struct {
	createdCoins []Scroogecoin
}

func (t *CoinCreationTransaction) ConsumedCoins() []ConsumedCoin {
	return []ConsumedCoin{}
}

func (t *CoinCreationTransaction) CreatedCoins() []Scroogecoin {
	return t.createdCoins
}

func (t *CoinCreationTransaction) Hash() []byte {
	hash := make([]byte, 0)

	for _, coin := range t.createdCoins {
		hash = append(hash, coin.Hash()...)
	}

	return hash
}

type PaymentTransaction struct {
	consumedCoins []ConsumedCoin
	createdCoins []Scroogecoin
}

func (t *PaymentTransaction) ConsumedCoins() []ConsumedCoin {
	return t.consumedCoins
}

func (t *PaymentTransaction) CreatedCoins() []Scroogecoin {
	return t.createdCoins
}

func (t *PaymentTransaction) Hash() []byte {
	hash := make([]byte, 0)

	for _, consumedCoin := range t.consumedCoins {
		hash = append(consumedCoin.hash())
	}

	for _, coin := range t.createdCoins {
		hash = append(hash, coin.Hash()...)
	}

	return hash
}