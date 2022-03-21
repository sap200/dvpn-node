package webapp

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/sap200/dvpn-node/client"
	"github.com/sap200/dvpn-node/server"
	"github.com/sap200/dvpn-node/utils"
	"github.com/sap200/vineyard/x/vineyard/types"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

var accountName string
var queryNode string
var loggerFile string
var accountUserName string
var isServerRunning bool
var hasConnection bool

// handle has arguments writer and reader (writer is the interface which has write function)
func handle(w http.ResponseWriter, req *http.Request) {

	// to list all the nodes
	nodeArr, err := client.QueryAll(queryNode)
	if err != nil {
		log.Fatalln(err)
	}

	// to read the logfile created in NewApp()
	a, err := utils.ReadFile(loggerFile)
	if err != nil {
		log.Println(err)
	}

	// struct declaration with details to be used and taken from index.html
	type d struct {
		Array   []types.Node
		Account string
		Logs    string
		Status  bool
		Session bool
	}

	// creation of x type struct
	x := d{
		Array:   nodeArr,
		Account: accountName,
		Logs:    string(a),
		Status:  isServerRunning,
		Session: hasConnection,
	}

	// parsing index.html template so that from here values can reach there
	parsedTemplate, _ := template.ParseFiles("./webapp/templates/index.html")
	// here we execute the parsing function with w as http response writer
	err = parsedTemplate.Execute(w, x)
	if err != nil {
		log.Println(err)
	}
}

func handleServer(w http.ResponseWriter, req *http.Request) {
	// used to read and request data filled into the html file using reader function req
	kh := req.FormValue("keyhome")
	rm := req.FormValue("remote")
	prt := req.FormValue("port")
	// used to strip
	kh = strings.Trim(kh, " ")
	rm = strings.Trim(rm, " ")
	prt = strings.Trim(prt, " ")

	// creating new RPC cosmos client
	// context.background is to run till the main program ends
	cc, err := cosmosclient.New(context.Background(),
		cosmosclient.WithNodeAddress(rm),
		cosmosclient.WithHome(kh),
	)
	if err != nil {
		log.Println(err)
	}

	isServerRunning = true

	// launched a zovino server goroutine
	go server.LaunchServer(cc, accountUserName, prt)

	nodeArr, err := client.QueryAll(queryNode)
	if err != nil {
		log.Fatalln(err)
	}

	a, err := utils.ReadFile(loggerFile)
	if err != nil {
		log.Println(err)
	}

	type d struct {
		Array   []types.Node
		Account string
		Logs    string
		Status  bool
		Session bool
	}

	x := d{
		Array:   nodeArr,
		Account: accountName,
		Logs:    string(a),
		Status:  isServerRunning,
		Session: hasConnection,
	}

	parsedTemplate, _ := template.ParseFiles("./webapp/templates/index.html")
	err = parsedTemplate.Execute(w, x)
	if err != nil {
		log.Println(err)
	}
}

func handleConnection(w http.ResponseWriter, req *http.Request) {

	kHome := req.FormValue("keyhome")
	remote := req.FormValue("remote")
	port := req.FormValue("port")
	ipaddr := req.FormValue("ip")
	// trim them
	kHome = strings.Trim(kHome, " ")
	remote = strings.Trim(remote, " ")
	port = strings.Trim(port, " ")
	ipaddr = strings.Trim(ipaddr, " ")

	// make a cosmos client
	cc, err := cosmosclient.New(context.Background(),
		cosmosclient.WithNodeAddress(remote),
		cosmosclient.WithHome(kHome),
	)

	if err != nil {
		log.Println(err)
	}

	// used to show the cosmos address for the given keyname
	add, err := cc.Address(accountUserName)
	if err != nil {
		log.Fatalln(err)
	}

	// generating a private public key pair for handshake
	// check handshake
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	c := client.NewClient(*privKey, ipaddr+":"+port, add.String())
	go c.Connect()

	hasConnection = true
	// Now the connection is formed and now is time to send output
	nodeArr, err := client.QueryAll(queryNode)
	if err != nil {
		log.Fatalln(err)
	}

	a, err := utils.ReadFile(loggerFile)
	if err != nil {
		log.Println(err)
	}

	type d struct {
		Array   []types.Node
		Account string
		Logs    string
		Status  bool
		Session bool
	}

	x := d{
		Array:   nodeArr,
		Account: accountName,
		Logs:    string(a),
		Status:  isServerRunning,
		Session: hasConnection,
	}

	parsedTemplate, _ := template.ParseFiles("./webapp/templates/index.html")
	err = parsedTemplate.Execute(w, x)
	if err != nil {
		log.Println(err)
	}

}

// NewApp makes a new HTTP app
// This is frontend for our VPN
func NewApp(port, accName, accUserName, qNode, logFile string) {

	// starts here

	// initialize variables
	accountName = accName
	queryNode = qNode
	loggerFile = logFile
	accountUserName = accUserName
	isServerRunning = false
	hasConnection = false

	// available handle functions "/" is the path

	// the router
	http.HandleFunc("/", handle)

	// automatically redirected to this per tab
	http.HandleFunc("/launcher", handleServer)
	http.HandleFunc("/connect", handleConnection)

	// this is used to start up and running a http servere at the specified port
	// (fatal ln is to log and return if there is any error)
	// the server
	log.Fatalln(http.ListenAndServe("localhost:"+port, nil))
}
