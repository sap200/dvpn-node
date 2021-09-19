package utils

import (
	"strings"
	"testing"
)

// test execute system command
func TestExecuteSystemCommand(t *testing.T) {
	command1 := []string{"which", "bash"}
	out, _ := executeSystemCommand(command1)

	if !strings.Contains(out.String(), "/usr/bin/bash") {
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

// test generation of random aes key
func TestGenerateRandomAESKey(t *testing.T) {
	k1 := GenerateRandomAESKey()
	k2 := GenerateRandomAESKey()

	if string(k1) == string(k2) {
		t.Fatalf("got k1 %s, got k2 %s", k1, k2)
	}
}

// test encryption, encryption may not be same always because of use of rand reader
func TestEncryptAES(t *testing.T) {
	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	text := "Hello world"
	se := EncryptAES(key, []byte(text))

	if len(se) != 79 {
		t.Fatalf("Expected %d, got %d", 79, len(se))
	}
}

// test decrypt aes
func TestDecryptAES(t *testing.T) {
	key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	texts := []string{"Hello world", "I am not good", "THis is testing", "Hello ! I am not here"}
	res := make([]string, len(texts))
	for i := 0; i < len(texts); i++ {
		se := EncryptAES(key, []byte(texts[i]))
		pt := DecryptAES(key, se)
		res[i] = string(pt)
	}

	for i := 0; i < len(texts); i++ {
		if res[i] != texts[i] {
			t.Fatalf("Expected %s, got %s", texts[i], res[i])
		}
	}

}
