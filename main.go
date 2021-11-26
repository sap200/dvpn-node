// Package main
package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sap200/dvpn-node/client"
	"github.com/sap200/dvpn-node/server"
	"github.com/sap200/dvpn-node/utils"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
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
	// take command line args for node tpye
	//nodeType := flag.String("node", "server", "the node type: server or client")
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	seed := serverCmd.String("seed", "http://localhost:26657", "node of a blockchain to connect to")
	accountName := serverCmd.String("account", "", "the account under which server should be registered")

	//query commands
	queryCmd := flag.NewFlagSet("list-nodes", flag.ExitOnError)
	qNode := queryCmd.String("seed", "http://localhost:26657", "query node address")

	// connect commands
	connectCmd := flag.NewFlagSet("connect", flag.ExitOnError)
	nodeID := connectCmd.String("ip", "", "ip address of the vpn node")
	accountAddress := connectCmd.String("account", "", "cosmos account name")

	if len(os.Args) < 2 {
		panic("please provide argument to be server or client")
	}

	switch os.Args[1] {
	case "server":
		serverCmd.Parse(os.Args[2:])

		fmt.Println("Here", *seed, "AccountName", *accountName)

		if *accountName == "" {
			os.Exit(1)
		}

		// installed := utils.InstallOpenvpn()
		// if !installed {
		// 	panic("unable to install openvpn")
		// }

		// launch server
		// make a cosmos client
		cc, err := cosmosclient.New(context.Background(), cosmosclient.WithNodeAddress(*seed))
		if err != nil {
			log.Fatalln(err)
		}

		utils.PrintServer()
		server.LaunchServer(cc, *accountName)

	case "list-nodes":
		queryCmd.Parse(os.Args[2:])

		nodeArr, err := client.QueryAll(*qNode)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println()
		for _, v := range nodeArr {
			fmt.Println("ID:", v.Index, "\tIP:", v.Address, "\tLocation: ", v.Location)
			fmt.Println()
		}

	case "connect":
		connectCmd.Parse(os.Args[2:])
		if *nodeID == "" {
			os.Exit(1)
		}

		if *accountAddress == "" {
			os.Exit(1)
		}

		utils.PrintClient()
		privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
		c := client.NewClient(*privKey, *nodeID+":8080", *accountAddress)
		c.Connect()

	}
}
