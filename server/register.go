package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/sap200/vineyard/x/vineyard/types"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

// Details stores details
type Details struct {
	IP       string `json:"IP"`
	City     string `json:"City"`
	Region   string `json:"Region"`
	Country  string `json:"Country"`
	Loc      string `json:"Loc"`
	Org      string `json:"Org"`
	Postal   string `json:"Postal"`
	Timezone string `json:"Timezone"`
	Readme   string `json:"Readme"`
}

func getDetails() Details {

	resp, err := http.Get(IPLink)
	if err != nil {
		log.Fatalln(err)
	}

	bod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var d Details
	err = json.Unmarshal(bod, &d)
	if err != nil {
		log.Fatalln(err)
	}

	return d
}

func registerNode(cc cosmosclient.Client, accountName string) {

	fmt.Println(cc)

	//accountName = "bob"
	// get account address from account name
	accountAddress, err := cc.Address(accountName)
	if err != nil {
		log.Fatalln(err)
	}

	// index of unique identity
	id := uuid.New()

	// detail of the Node
	detail := getDetails()

	// create a new message of type create Node
	msg := types.NewMsgCreateNode(
		accountAddress.String(),
		id.String(),
		detail.IP,
		detail.Region+", "+Iso2ToCountry[detail.Country],
		"", //future bandwidth
		"", // future uid or empty field
	)

	// Broadcast the Txs
	txResp, err := cc.BroadcastTx(accountName, msg)
	if err != nil {
		log.Fatalln(err)
	}

	// if successful, assign the id to UUIDOfServer
	// It will be useful while deleting the txs
	UUIDOfServer = id.String()
	// log the registry txs
	log.Println(txResp)

}
