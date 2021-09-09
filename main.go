// Package main
package main

import (
	"github.com/sap200/dvpn-node/utils"
	"github.com/sap200/dvpn-node/server"
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
	utils.InstallOpenvpn()

	// launch server
	server.LaunchServer()
	
}