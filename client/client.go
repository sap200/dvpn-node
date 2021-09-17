// Package client gives an option to connect to other vpn servers
package client

import "crypto/rsa"

// Client contains general information about Client
type Client struct {
	PrivateKey    rsa.PrivateKey `json:"private_key"`
	ServerAddress string         `json:"server_address"`
}

// NewClient creates a new Client and connects to the server
func NewClient(privateKey rsa.PrivateKey, serverAddress string) Client {
	c := Client{
		PrivateKey:    privateKey,
		ServerAddress: serverAddress,
	}

	return c
}

// Connect method lets the client handshake and make a openvpn tunnel to the given server
func (c Client) Connect() {
	// installs openvpn

	// client generates new private and public key of its own and sets it

	// client sends TCP syn packet to server

	// client gets back ack packet decodes it and reads the message

	// verifies the message and writes the configuration to a file

	// once done it executes open-vpn command to connect to the respective server

}
