package server

import (
	"crypto/rsa"

	"github.com/sap200/dvpn-node/packets"
)

// Storage stores the public keys of the incoming connection
var Storage map[string]packets.SynPacket

// MyPrivateKey contains my own private key for secure connection
var MyPrivateKey rsa.PrivateKey

// MyPublicKey contains public key of the server
var MyPublicKey rsa.PublicKey

// PATHPREFIX is a constant denoting where the clients are stored
const PATHPREFIX = "/root/"

// InitStore initializes the store keuy
func InitStore() map[string]rsa.PublicKey {
	return map[string]rsa.PublicKey{}
}
