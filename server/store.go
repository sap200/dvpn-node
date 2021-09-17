package server

import (
	"crypto/rsa"
)

// Storage stores the public keys of the incoming connection
var Storage map[string]rsa.PublicKey

// MyPrivateKey contains my own private key for secure connection
var MyPrivateKey rsa.PrivateKey

// MyPublicKey contains public key of the server
var MyPublicKey rsa.PublicKey

// InitStore initializes the store keuy
func InitStore() map[string]rsa.PublicKey {
	return map[string]rsa.PublicKey{}
}
