// Package client gives an option to connect to other vpn servers
package client

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os/exec"

	"github.com/sap200/dvpn-node/packets"
	"github.com/sap200/dvpn-node/utils"
)

// Client contains general information about Client
type Client struct {
	PrivateKey    rsa.PrivateKey `json:"private_key"`
	ServerAddress string         `json:"server_address"`
	CosmosAddress string         `json:"cosmos_address"`
}

// NewClient creates a new Client and connects to the server
func NewClient(privateKey rsa.PrivateKey, serverAddress, cosmosAddress string) Client {
	c := Client{
		PrivateKey:    privateKey,
		ServerAddress: serverAddress,
		CosmosAddress: cosmosAddress,
	}

	return c
}

// Connect method lets the client handshake and make a openvpn tunnel to the given server
func (c Client) Connect() {
	// assume openvpn is already installed
	// client sends TCP syn packet to server
	synPacket := packets.NewSynPacket(c.PrivateKey.PublicKey, c.CosmosAddress)
	synPackString, err := synPacket.Marshall()
	check(err)

	// send synPacket to the connection
	con, err := net.Dial("tcp", c.ServerAddress)
	check(err)

	_, err = io.WriteString(con, synPackString)
	check(err)

	// client gets back ack packet decodes it and reads the message
	b, err := bufio.NewReader(con).ReadString('\n')
	check(err)

	// decode the message
	var ackPack packets.AckPacket
	err = json.Unmarshal([]byte(b), &ackPack)
	check(err)

	var msgPacket packets.MsgPacket
	// unmarshal message
	err = json.Unmarshal([]byte(ackPack.Message), &msgPacket)
	check(err)

	// verify the signature of the message
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, &c.PrivateKey, msgPacket.Cipher)
	check(err)

	hasher := sha512.New()
	hasher.Write(plainText)
	hash := hasher.Sum(nil)

	// verifies the message and writes the configuration to a file
	err = rsa.VerifyPKCS1v15(&ServerPublicKey, crypto.SHA512, hash, msgPacket.Signature)
	check(err)

	// once done it executes open-vpn command to connect to the respective server
	err = utils.WriteFile("./openvpn-connection.ovpn", plainText)
	check(err)

	// connect to openvpn using the file
	out, inErr := executeSystemCommand([]string{"openvpn", "./openvpn-connection.ovpn"})
	fmt.Println(inErr.String())
	fmt.Println(out.String())
}

// checks if the error occured
// if err is not nil
// then panics and quits out of the program
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// executes system command specifically
// designed to run openvpn command
// in this client
func executeSystemCommand(command []string) (bytes.Buffer, bytes.Buffer) {
	var out bytes.Buffer
	var err bytes.Buffer // modified
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = &out
	cmd.Stderr = &err // modified
	cmd.Run()

	return out, err
}
