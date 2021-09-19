// Package main
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"flag"

	"github.com/sap200/dvpn-node/client"
	"github.com/sap200/dvpn-node/server"
	"github.com/sap200/dvpn-node/utils"
)

// Aashin's part

func main() {

	// TODO: take the command line argument
	// for now no command line args
	// just call the install openvpn function
	// check the status if openvpn-server is active by doing sudo service openvpn-service | grep status
	// if status is active then print everything done successfully
	// else print something went wrong...

	// install and start openvpn as a service
	nodeType := flag.String("node", "server", "the node type: server or client")
	flag.Parse()
	if *nodeType == "server" {
		installed := utils.InstallOpenvpn()
		if !installed {
			panic("unable to install openvpn")
		}

		// launch server
		server.LaunchServer()

	} else if *nodeType == "client" {
		privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
		c := client.NewClient(*privKey, "localhost:8080", "cosmos11abcxergtyds")
		c.Connect()
	}

}
