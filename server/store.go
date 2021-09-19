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

// EXTENSION is a constant for openvpn file which is .ovpn
const EXTENSION = ".ovpn"

// AesKey is used to encrypt and decrypt message
var AesKey []byte

// InitStore initializes the store keuy
func InitStore() map[string]packets.SynPacket {
	return map[string]packets.SynPacket{}
}
