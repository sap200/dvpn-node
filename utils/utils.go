package utils

// Benher's part

// install openvpn-server and configure it also run it as a service
// bash command "./scripts/openvpn-install.sh"
func InstallOpenvpn() {
	// TODO: Install openvpn log the output
}

// add a client to the openvpn with the given client name
// bash command "bash ./scripts/add-client.sh clientName"
func AddClient(clientName string) bool {
	// TODO: add a new client with specified name...
	return false
}

// remove a client from openvpn access list
// bash command "bash "./scripts/remove-client.sh clientName"
func RevokeClient(clientName string) bool {
	// TODO: revoke client access or remove client to access openvpn
	return false
}