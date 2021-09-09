// package server
// launches a server to listen to incoming requests and deliver the responses accordingly

// this done in handleConnection
// in case of intiated message it sends the pubkey
// in case of terminated message it revokes the client access and deletes the public key.
package server

import (
	"fmt"
	"github.com/sap200/dvpn-node/utils"
	"net"
) 

func LaunchServer() {
	// log the server start
	fmt.Printf("Node started at port %s\n", utils.PORT)

	ln, err := net.Listen("tcp", ":" + utils.PORT)
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
	// TODO: handle handshake
	// terminate the connection
}