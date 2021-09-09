package utils

import (
	"testing"
	"strings"

)

// test execute system command
func TestExecuteSystemCommand(t *testing.T) {
	command1 := []string{"which", "bash"}
	out, _ := executeSystemCommand(command1)

	if !strings.Contains(out.String(),  "/usr/bin/bash") {
		t.Fatalf("Expected %s, got %s", "/usr/bin/bash", out.String())
	} 
}

// test install openvpn
func TestInstallOpenvpn(t *testing.T) {
	command := []string{"bash", "." + OPENVPN_INSTALL_PATH}
	out, _ := executeSystemCommand(command)
	if !strings.Contains(out.String(), OPENVPN_INSTALL_SUCCESS_FLAG1) &&
	!strings.Contains(out.String(), OPENVPN_INSTALL_SUCCESS_FLAG2) {
		t.Fatalf("Expected %s, got %s", OPENVPN_INSTALL_SUCCESS_FLAG1, out.String())
	} 
}

// test add client
func TestAddClient(t *testing.T) {
	command := []string{"bash", "." + OPENVPN_ADD_CLIENT_PATH, "axgQQDERllmnp"}
	executeSystemCommand(command)

	out, _ := executeSystemCommand([]string{"ls", "/root"})
	if !strings.Contains(out.String(), "axgQQDERllmnp") {
		t.Fatalf("Expected %s, got %s", out.String(), "None")
	}
}

// test revoke client
func TestRevokeClient(t *testing.T) {
	command := []string{"bash", "." + OPENVPN_REMOVE_CLIENT_PATH, "axgQQDERllmnp"}
	out, _ := executeSystemCommand(command)
	if !strings.Contains(out.String(), "revoked!") {
		t.Fatalf("Expected %s, got %s", "revoked!", out.String())
	}
}