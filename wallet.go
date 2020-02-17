package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewWallet() *Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		panic(err)
	}

	return &Wallet{publicKey: &privateKey.PublicKey, privateKey: privateKey}
}

func (w *Wallet) sign(message []byte) (r *big.Int, s *big.Int) {
	rng := rand.Reader
	hash := sha256.Sum256(message)

	r, s, err := ecdsa.Sign(rng, w.privateKey, hash[:])

	if err != nil {
		panic(err)
	}

	return
}

func (w *Wallet) verifySignature(publicKey *ecdsa.PublicKey, r *big.Int, s *big.Int, message []byte) bool {
	hash := sha256.Sum256(message)

	return ecdsa.Verify(publicKey, hash[:], r, s)
}

func (w *Wallet) GetWalletId() *ecdsa.PublicKey {
	return w.publicKey
}
