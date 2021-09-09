package server

// launches a server to listen to incoming requests and deliver the responses accordingly

// this done in handleConnection
// in case of intiated message it sends the pubkey
// in case of terminated message it revokes the client access and deletes the public key.

// Aashin's part..
func LaunchServer() {

ln, err := net.Listen("tcp", ":8080")
if err != nil {
	panic(err)
}
for {
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	go HandleConnection(conn)
}

}

func HandleConnection() {
	
}