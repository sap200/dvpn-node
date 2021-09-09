package utils

import (
    "bytes"
    "os/exec"
)
//execute linux commands
func executeSystemCommand(command []string) bool {
    var out bytes.Buffer
    var err bytes.Buffer // modified

    cmd := exec.Command(...command)
    cmd.Stdout = &out
    cmd.Stderr = &err // modified


    if len(err.String()) != 0 {
        return false
    } else {
        return true
    }
}
// Benher's part

// install openvpn-server and configure it also run it as a service
// bash command "./scripts/openvpn-install.sh"
func InstallOpenvpn() bool {
	command := []string{"bash", "../scripts/openvpn-install.sh"}
    	return executeSystemCommand(command)
}

// add a client to the openvpn with the given client name
// bash command "bash ./scripts/add-client.sh clientName"
func AddClient(clientName string) bool {
	command := []string{"bash", "../scripts/add-client.sh", clientName}
    	return executeSystemCommand(command)
}

// remove a client from openvpn access list
// bash command "bash "./scripts/remove-client.sh clientName"
func RevokeClient(clientName string) bool {
	command := []string{"bash", "../scripts/remove-client.sh", clientName}
    	return executeSystemCommand(command)
}
