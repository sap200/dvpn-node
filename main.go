// Package main
package main

import "utils"
import "server"
// Aashin's part

func main() {


	utils.InstallOpenvpn()

	server.LaunchServer()
	// TODO: take the command line argument 
	// for now no command line args
	// just call the install openvpn function
	// check the status if openvpn-server is active by doing sudo service openvpn-service | grep status 
	// if status is active then print everything done successfully
	// else print something went wrong...
}