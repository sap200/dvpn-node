package server

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/sap200/dvpn-node/packets"
)

// Storage stores the public keys of the incoming connection
var Storage map[string]packets.SynPacket

// MyPrivateKey contains my own private key for secure connection
var MyPrivateKey rsa.PrivateKey

// MyPublicKey contains public key of the server
var MyPublicKey rsa.PublicKey

// PATHPREFIX is a constant denoting where the clients are stored
const PATHPREFIX = "/root/"

// EXTENSION is a constant for openvpn file which is .ovpn
const EXTENSION = ".ovpn"

// AesKey is used to encrypt and decrypt message
var AesKey []byte

// IPLink is a variable that contains the link to make request to get ip address
var IPLink = "http://ipinfo.io"

// Iso2ToCountry contains Iso2 code
var Iso2ToCountry = map[string]string{}

// iso2Link is the link with json file containing iso2 to country code
var iso2Link = "https://datahub.io/core/country-list/r/0.json"

// UUIDOfServer is the unique identity to identify the server
var UUIDOfServer string

// InitStore initializes the store keuy
func InitStore() map[string]packets.SynPacket {
	return map[string]packets.SynPacket{}
}

// GetIP gets the Ip of the addresss string
func GetIP(addr string) string {
	return strings.Split(addr, ":")[0]
}

// CountryCode is a struct containing country code and its corresponding full name
type CountryCode struct {
	Code string
	Name string
}

func initMap() {
	res, err := http.Get(iso2Link)
	if err != nil {
		log.Fatalln(err)
	}

	bod, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var lst []CountryCode

	err = json.Unmarshal(bod, &lst)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range lst {
		Iso2ToCountry[v.Code] = v.Name
	}

}
