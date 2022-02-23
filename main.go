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
	"github.com/sap200/dvpn-node/parser"
	"github.com/sap200/dvpn-node/server"
	"github.com/sap200/dvpn-node/utils"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

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
	//------------------- server command-------------------
	serverCmd := flag.NewFlagSet("server", flag.ContinueOnError)
	cfileName := serverCmd.String("config", "", "Configuration File")
	//seed := serverCmd.String("seed", utils.DEFAULT_BLOCKCHAIN_LINK, "node of a blockchain to connect to")
	//accountName := serverCmd.String("account", "", "the account under which server should be registered")

	//query commands
	queryCmd := flag.NewFlagSet("list-nodes", flag.ContinueOnError)
	qNode := queryCmd.String("seed", utils.DEFAULT_BLOCKCHAIN_LINK, "query node address")

	// connect commands
	connectCmd := flag.NewFlagSet("connect", flag.ContinueOnError)
	confile := connectCmd.String("config", "", "Configuration file")
	//nodeID := connectCmd.String("ip", "", "ip address of the vpn node")
	//accountName1 := connectCmd.String("account", "", "cosmos account name")

	flag.Parse()

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		serverCmd.Parse(os.Args[2:])

		//fmt.Println("Here", *seed, "AccountName", *accountName)

		if *cfileName == "" {
			serverCmd.PrintDefaults()
			os.Exit(1)
		}

		// check installation
		installed := utils.InstallOpenvpn()
		if !installed {
			panic("unable to install openvpn")
		}

		// Parse the config file
		cfg, err := parser.ParseServerConfig(*cfileName)
		if err != nil {
			panic(err)
		}

		// launch server
		// make a cosmos client
		cc, err := cosmosclient.New(context.Background(),
			cosmosclient.WithNodeAddress(cfg.Remote),
			cosmosclient.WithHome(cfg.KeyHome),
		)

		if err != nil {
			log.Fatalln(err)
		}

		d, e := cc.Address(cfg.Account)
		fmt.Println("Here: d", d)
		fmt.Println("Here: e", e)

		utils.PrintServer()
		server.LaunchServer(cc, cfg.Account, cfg.Port)

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
		if *confile == "" {
			connectCmd.PrintDefaults()
			os.Exit(1)
		}

		cfg, err := parser.ParseSessionConfig(*confile)
		if err != nil {
			panic(err)
		}

		cc, err := cosmosclient.New(context.Background(),
			cosmosclient.WithNodeAddress(cfg.Remote),
			cosmosclient.WithHome(cfg.KeyHome),
		)

		if err != nil {
			log.Fatalln(err)
		}

		add, err := cc.Address(cfg.Account)
		if err != nil {
			log.Fatalln(err)
		}

		//fmt.Println("Address", add.String())

		utils.PrintClient()
		privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
		c := client.NewClient(*privKey, cfg.IPAddr+":"+cfg.Port, add.String())
		c.Connect()

	default:
		log.Println("Invalid command")
		printUsage()
		os.Exit(1)

	}
}

func printUsage() {
	fmt.Println("usage:")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("server\n\t-for launching a server")
	fmt.Println("list-nodes\n\t-for listing available nodes")
	fmt.Println("connect\n\t-for connecting to available nodes")
}
