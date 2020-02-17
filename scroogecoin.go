package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"fmt"
)

type Scroogecoin struct {
	value    float64
	walletId *ecdsa.PublicKey
	transactionId []byte
	coinNum int
}

func NewCoin(value float64, walletId *ecdsa.PublicKey, transactionId []byte, coinNum int) Scroogecoin {
	if value <= 0 {
		panic("value must be greater than 0.")
	}

	if walletId == nil {
		panic("walledId cannot be nil")
	}

	return Scroogecoin{value: value, walletId: walletId, transactionId: transactionId, coinNum: coinNum}
}

func (s *Scroogecoin) Equals(s1 *Scroogecoin) bool {
	if len(s.transactionId) != len(s1.transactionId) {
		return false
	}

	for b := range s.transactionId {
		if s.transactionId[b] != s1.transactionId[b] {
			return false
		}
	}

	return s.coinNum == s1.coinNum
}

func (s *Scroogecoin) ToString() string {
	var str string = ""
	return "Num: " + str + ", Value: " + fmt.Sprintf("%f", s.value) + ", Wallet id: " + fmt.Sprintf("%d", s.walletId.Y)
}

func (s *Scroogecoin) Hash() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, s.value)
	buf.Write(s.publicKeyBytes())
	binary.Write(buf, binary.BigEndian, s.coinNum)

	bytes := append(buf.Bytes(), s.transactionId...)
	hash := sha256.Sum256(bytes)

	return hash[:]
}

func (s *Scroogecoin) publicKeyBytes() []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(s.walletId)

	if err != nil {
		panic(err)
	}

	return pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})
}
