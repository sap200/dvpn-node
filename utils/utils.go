package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
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

// InstallOpenvpn install openvpn-server and configure it also run it as a service
// bash command "./scripts/openvpn-install.sh"
func InstallOpenvpn() bool {
	command := []string{"bash", OPENVPN_INSTALL_PATH}
	out, _ := executeSystemCommand(command)
	fmt.Println(out.String())
	if strings.Contains(out.String(), OPENVPN_INSTALL_SUCCESS_FLAG1) ||
		strings.Contains(out.String(), OPENVPN_INSTALL_SUCCESS_FLAG2) {
		return true
	}
	return false
}

// AddClient add a client to the openvpn with the given client name
// bash command "bash ./scripts/add-client.sh clientName"
func AddClient(clientName string) bool {
	command := []string{"bash", OPENVPN_ADD_CLIENT_PATH, clientName}
	out, _ := executeSystemCommand(command)
	if strings.Contains(out.String(), OPENVPN_ADD_CLIENT_SUCCESS_FLAG) {
		return true
	}
	return false
}

// RevokeClient remove a client from openvpn access list
// bash command "bash "./scripts/remove-client.sh clientName"
func RevokeClient(clientName string) bool {
	command := []string{"bash", OPENVPN_REMOVE_CLIENT_PATH, clientName}
	out, _ := executeSystemCommand(command)
	if strings.Contains(out.String(), OPENVPN_REVOKE_CLIENT_SUCCESS_FLAG) {
		return true
	}
	return false

}

// ReadFile reads a file
func ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

// WriteFile writes to a file
func WriteFile(filePath string, data []byte) error {
	f, err := os.Create(filePath)
	defer f.Close()

	if err != nil {
		return err
	}

	if err := os.Chmod(filePath, 0666); err != nil {
		panic(err)
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// StopOpenvpn stops the openvpn service by executing a system command
func StopOpenvpn() {
	cmd := exec.Command("systemctl", "stop", "openvpn")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
