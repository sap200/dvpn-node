package utils

import (
	"bytes"
	"os/exec"
	"strings"
	"fmt"
)
//execute linux commands
func executeSystemCommand(command []string) (bytes.Buffer, bytes.Buffer) {
    var out bytes.Buffer
    var err bytes.Buffer // modified
	cmd := exec.Command(command[0], command[1:]...)
    cmd.Stdout = &out
	cmd.Stderr = &err // modified
	cmd.Run()

	return out, err
}

// install openvpn-server and configure it also run it as a service
// bash command "./scripts/openvpn-install.sh"
func InstallOpenvpn() bool {
	command := []string{"bash", OPENVPN_INSTALL_PATH}
	out, _ := executeSystemCommand(command)
	fmt.Println(out.String())
	if strings.Contains(out.String(), OPENVPN_INSTALL_SUCCESS_FLAG1) || 
	strings.Contains(out.String(), OPENVPN_INSTALL_SUCCESS_FLAG2) {
		return true
	} else {
		return false
	}
}

// add a client to the openvpn with the given client name
// bash command "bash ./scripts/add-client.sh clientName"
func AddClient(clientName string) bool {
	command := []string{"bash", OPENVPN_ADD_CLIENT_PATH, clientName}
	out, _ := executeSystemCommand(command)
	if strings.Contains(out.String(), OPENVPN_ADD_CLIENT_SUCCESS_FLAG) {
		return true
	} else {
		return false
	}
}

// remove a client from openvpn access list
// bash command "bash "./scripts/remove-client.sh clientName"
func RevokeClient(clientName string) bool {
	command := []string{"bash", OPENVPN_REMOVE_CLIENT_PATH, clientName}
	out, _ := executeSystemCommand(command)
	if strings.Contains(out.String(), OPENVPN_REVOKE_CLIENT_SUCCESS_FLAG) {
		return true
	} else {
		return false
	}
}
