// Package server launches a server to listen to incoming requests and deliver the responses accordingly
// this done in handleConnection
// in case of intiated message it sends the pubkey
// in case of terminated message it revokes the client access and deletes the public key.
package server

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/sap200/dvpn-node/packets"
	"github.com/sap200/dvpn-node/utils"
)

func initVars() {
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	MyPrivateKey = *pk
	MyPublicKey = pk.PublicKey
	AesKey = utils.GenerateRandomAESKey()
	Storage = InitStore()
}

// LaunchServer launches the server
func LaunchServer() {
	// init the variables
	initVars()
	// log the server start
	fmt.Printf("Node started at port %s\n", utils.PORT)

	// on press of ctrl + c, do the basic cleanup before exiting the server
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go cleanup(sigs)

	// start the listener
	ln, err := net.Listen("tcp", ":"+utils.PORT)
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
	// terminate the connection...
	fmt.Println("processing request from", conn.RemoteAddr().String())
	defer conn.Close()

	// receive synPacket
	b, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		// unable to read SynPacket
		fmt.Println("Unable to read incoming handshake from", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}
	// decode synPacket
	var synPacket packets.SynPacket
	err = json.Unmarshal([]byte(b), &synPacket)
	if err != nil {
		fmt.Println("Unable to unmarshall SynPacket form", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}

	if synPacket.Message == utils.TERMINATE {
		fmt.Println("Received termination command from ", conn.RemoteAddr().String())
		// delete the client from the storage
		pack, ok := Storage[GetIP(conn.RemoteAddr().String())]
		stat := packets.AckFail
		if ok {
			// Revoke client access
			utils.RevokeClient(pack.Message)
			// delete it from the storage
			delete(Storage, GetIP(conn.RemoteAddr().String()))
			// set the ack status to success
			stat = packets.AckSuccess
		}
		// make an ack packet and send back to the client
		ackPacket := packets.NewAckPacket(int64(stat), rsa.PublicKey{}, "terminated connection")
		ackMsg, err := ackPacket.Marshall()
		if err != nil {
			fail(conn, err.Error())
		}

		fmt.Println(ackPacket)
		_, err = io.WriteString(conn, ackMsg)
		if err != nil {
			fmt.Println("Unable to write to connection")
		}

		if stat == packets.AckSuccess {
			fmt.Println("Connection termination was successful..")
		} else {
			fmt.Println("Unsuccessful connection termination")
		}

		return
	}

	// store in corresponding ip its public key
	Storage[GetIP(conn.RemoteAddr().String())] = synPacket
	// create a client
	clientAdded := utils.AddClient(synPacket.Message)
	if !clientAdded {
		fmt.Println("Unable to add client", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}

	// encrypt the message with public key
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, &synPacket.PubKey, AesKey)
	if err != nil {
		fmt.Println(err, "for", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}

	// sign with own private key
	hasher := sha512.New()
	hasher.Write(AesKey)
	hash := hasher.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, &MyPrivateKey, crypto.SHA512, hash)
	if err != nil {
		fmt.Println(err, "for", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}

	msgPacket := packets.NewMsgPacket(cipher, sign)
	msgString, err := msgPacket.Marshall()
	if err != nil {
		fmt.Println(err, "for", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}
	// send a AckPacket packet containing openvpn file
	ackPacket := packets.NewAckPacket(packets.AckSuccess, MyPublicKey, msgString)
	ackString, err := ackPacket.Marshall()
	if err != nil {
		fmt.Println(err, "for", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}

	// send your public key to the client
	_, err = io.WriteString(conn, ackString)
	if err != nil {
		fmt.Println(err, "for", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}

	// done acknowledgement..
	// receive new acknowledgement from client representing aes key received
	b, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err, "for", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}
	var ackPack packets.AckPacket
	err = json.Unmarshal([]byte(b), &ackPack)
	if err != nil {
		fmt.Println(err, "for", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}
	if ackPack.AckStatus != packets.AckSuccess {
		fmt.Println(err, "for", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}

	//fmt.Println(ackPack)

	// got the aes key now...
	// read the openvpn file and encrypt using aes key
	data, err := utils.ReadFile(PATHPREFIX + synPacket.Message + EXTENSION)
	if err != nil {
		fmt.Println(err)
		fail(conn, "Handshake Failed")
		return
	}

	// now encrypt using AES data and send it back
	encrypted := utils.EncryptAES(AesKey, data)

	// transfer aes encrypted file
	_, err = io.WriteString(conn, encrypted)
	if err != nil {
		fmt.Println(err, "for", conn.RemoteAddr())
		fail(conn, "Handshake Failed")
		return
	}

	fmt.Println("Establishing secure tunnel with", conn.RemoteAddr().String())
	//fmt.Println(encrypted)
	// aes encrypted file transferred
	// rest of the things are taken care by openvpn server
	// .... finish of server

}

// When the handshake fails fail writes the acknowledgement fail packet to the connection.
func fail(conn net.Conn, message string) {
	var x rsa.PublicKey
	ackPacket := packets.NewAckPacket(packets.AckFail, x, message)
	bs, err := ackPacket.Marshall()
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, bs)
	conn.Close()
	fmt.Println("Handshake Failed closed connection")
}

// Execute system command
func executeSystemCommand(command []string) bool {
	var out bytes.Buffer
	var err bytes.Buffer // modified
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = &out
	cmd.Stderr = &err // modified
	cmd.Run()

	if err.String() == "" {
		return false
	}

	return true
}

// cleanup function cleans up everything before exiting the server
func cleanup(sigs chan os.Signal) {
	sig := <-sigs
	fmt.Println(sig, "Inside cleanup routine")
	// ----------------------------------------------------------------
	//TODO: deregister the server on cosmos blockchain, that it went down...
	// This is to be done, after the blockchain is ready

	// -----------------------------------------------------------------
	// cleanup all the existing users in the storage
	for _, pack := range Storage {
		fmt.Print("Revoking ", pack.Message)
		res := utils.RevokeClient(pack.Message)
		if res {
			fmt.Print(" ,Revoke Success")
		} else {
			fmt.Print(" ,Revoke Fail")
		}
		fmt.Println()
	}

	// stop the openvpn server
	utils.StopOpenvpn()
	// exit
	os.Exit(0)
}
