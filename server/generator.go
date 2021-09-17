package server

import (
	"crypto/rand"
	"crypto/rsa"
)

// NewGenerator generates public key
func NewGenerator() rsa.PrivateKey {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	return *privKey
}
