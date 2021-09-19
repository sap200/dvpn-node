package client

import "crypto/rsa"

// ServerPublicKey contains Public Key of the server
var ServerPublicKey rsa.PublicKey

// AesKey is the key received by server
// it is set and rest communications happens by AES encryption
var AesKey []byte
