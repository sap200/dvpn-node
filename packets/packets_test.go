package packets

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"testing"
)

// TestNewSynPacket tests SynPacket
func TestNewSynPacket(t *testing.T) {
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	synPacket := NewSynPacket(pk.PublicKey, "Public key of the server sent")
	if synPacket.PubKey != pk.PublicKey {
		t.Fatalf("Expected: %v, got: %v", pk.PublicKey, synPacket.PubKey)
	}

	if synPacket.Message != "Public key of the server sent" {
		t.Fatalf("Expected: Public key of the server sent, got: %s", synPacket.Message)
	}
}

// TestNewAckPacket tests SynPacket
func TestNewAckPacket(t *testing.T) {
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	ackPacket := NewAckPacket(1, pk.PublicKey, "Acknowledgement packet")
	if ackPacket.AckStatus != 1 {
		t.Fatalf("Expected: %d, got %d", 1, ackPacket.AckStatus)
	}

	if ackPacket.PubKey != pk.PublicKey {
		t.Fatalf("Expected: %v, got %v", pk.PublicKey, ackPacket.PubKey)
	}

	if ackPacket.Message != "Acknowledgement packet" {
		t.Fatalf("Expected: %s, got: %s", "Acknowledgement packet", ackPacket.Message)
	}
}

// TestNewMsgPacket tests MsgPacket
func TestMsgPacket(t *testing.T) {
	text := "Hello: How are you ??"
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	hasher := sha512.New()
	hasher.Write([]byte(text))
	hash := hasher.Sum(nil)

	// signature
	sign, err := rsa.SignPKCS1v15(rand.Reader, pk, crypto.SHA512, hash)
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, &pk.PublicKey, []byte(text))

	msgPacket := NewMsgPacket(cipher, sign)
	if string(msgPacket.Cipher) != string(cipher) {
		t.Fatalf("Expected: %v, got: %v", cipher, msgPacket.Cipher)
	}

	if string(msgPacket.Signature) != string(sign) {
		t.Fatalf("Expected: %v, got: %v", sign, msgPacket.Signature)
	}
}
