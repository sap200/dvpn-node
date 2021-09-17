// Package server launches a server to listen to incoming requests and deliver the responses accordingly
// this done in handleConnection
// in case of intiated message it sends the pubkey
// in case of terminated message it revokes the client access and deletes the public key.
package server

import (
	"bufio"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/sap200/dvpn-node/packets"
	"github.com/sap200/dvpn-node/utils"
)

// LaunchServer launches the server
func LaunchServer() {
	// log the server start
	fmt.Printf("Node started at port %s\n", utils.PORT)

	ln, err := net.Listen("tcp", ":"+utils.PORT)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// terminate the connection...
	defer conn.Close()

	// receive synPacket
	b, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		// unable to read SynPacket
		fmt.Println("Unable to read incoming handshake from", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}
	// decode synPacket
	var synPacket packets.SynPacket
	err = json.Unmarshal([]byte(b), &synPacket)
	if err != nil {
		fmt.Println("Unable to unmarshall SynPacket form", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}
	// store in corresponding ip its public key
	Storage[conn.RemoteAddr().String()] = synPacket.PubKey

	// encrypt the message with public key

	// sign with own private key
	// send a message packet containing openvpn file
	// send your public key to the client

}

// When the handshake fails fail writes the acknowledgement fail packet to the connection.
func fail(conn net.Conn, message string) {
	var x rsa.PublicKey
	ackPacket := packets.NewAckPacket(packets.AckFail, x, message)
	bs, err := ackPacket.Marshall()
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, bs)
	conn.Close()
	fmt.Println("Handshake Failed closed connection")
}
