package webapp

import (
	"html/template"
	"log"
	"net/http"

	"github.com/sap200/dvpn-node/client"
	"github.com/sap200/dvpn-node/utils"
	"github.com/sap200/vineyard/x/vineyard/types"
)

var accountName string
var queryNode string
var loggerFile string

func handle(w http.ResponseWriter, req *http.Request) {

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
	}

	x := d{
		Array:   nodeArr,
		Account: accountName,
		Logs:    string(a),
	}

	parsedTemplate, _ := template.ParseFiles("./webapp/templates/index.html")
	err = parsedTemplate.Execute(w, x)
	if err != nil {
		log.Println(err)
	}
}

// NewApp makes a new HTTP app
// This is frontend for our VPN
func NewApp(port, accName, qNode, logFile string) {
	accountName = accName
	queryNode = qNode
	loggerFile = logFile
	http.HandleFunc("/", handle)
	log.Fatalln(http.ListenAndServe("localhost:"+port, nil))
}
