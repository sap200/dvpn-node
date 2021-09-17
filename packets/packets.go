package packets

import (
	"crypto/rsa"
	"encoding/json"
)

// Packet defines a interface with a Marshall method
// that any types if implementing packet should implement
type Packet interface {
	Marshall()
}

// SynPacket is the packet
// that is sent by the client while establishing secure connection
type SynPacket struct {
	PubKey  rsa.PublicKey `json:"public_key"`
	Message string        `json:"message"`
}

// NewSynPacket creates a new packet of type SynPacket
func NewSynPacket(pubKey rsa.PublicKey, message string) SynPacket {
	a := SynPacket{
		PubKey:  pubKey,
		Message: message,
	}

	return a
}

// Marshall converts SynPacket to String to transfer using TCP protocol
func (p SynPacket) Marshall() (string, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(bs) + "\n", nil
}

// AckPacket contains the acknowledgement
// AckStatus defines the acknowledgement status i.e. failed or success
type AckPacket struct {
	AckStatus int64         `json:"ack_status"`
	PubKey    rsa.PublicKey `json:"public_key"`
	Message   string        `json:"message"`
}

// NewAckPacket creates a new Acknowledgement packet
func NewAckPacket(ackStatus int64, pubKey rsa.PublicKey, message string) AckPacket {
	a := AckPacket{
		AckStatus: ackStatus,
		PubKey:    pubKey,
		Message:   message,
	}

	return a
}

// Marshall converts AckPacket to string for TCP transfer
func (p AckPacket) Marshall() (string, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(bs) + "\n", nil
}

// MsgPacket contains the cipher and signature
// signature needs to be validated and cipher needs to be decrypted
type MsgPacket struct {
	Cipher    []byte `json:"cipher_text"`
	Signature []byte `json:"signature"`
}

// NewMsgPacket makes new msg packet
func NewMsgPacket(cipher []byte, signature []byte) MsgPacket {
	m := MsgPacket{
		Cipher:    cipher,
		Signature: signature,
	}

	return m
}

// Marshall converts AckPacket to string for TCP transfer
func (p MsgPacket) Marshall() (string, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(bs) + "\n", nil
}
