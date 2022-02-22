package parser

import (
	"os"
	"testing"
)

func TestParseServerConfig(t *testing.T) {

	name := "test_config.json"
	f, err := os.Create(name)
	if err != nil {
		t.Fatalf(err.Error())
	}

	d := `{
	"Account": "alice",
	"Remote": "http://localhost:26657",
	"KeyHome": "/home/sapta/vineyard",
	"Port": "5899"
}`

	f.Write([]byte(d))

	pc, err := ParseServerConfig(name)
	if err != nil {
		t.Fatalf(err.Error())
	}

	os.Remove(name)

	//fmt.Println(pc)

	if pc.Account != "alice" {
		t.Fatalf("Account mismatch: Expected %s, got %s\n", "alice", pc.Account)
	}

	if pc.Remote != "http://localhost:26657" {
		t.Fatalf("Remote mismatch: Expected %s, got %s\n", "http://localhost:26657", pc.Remote)
	}

	if pc.KeyHome != "/home/sapta/vineyard" {
		t.Fatalf("KeyHome mismatch: Expected %s, got %s\n", "/home/sapta/vineyard", pc.KeyHome)
	}

	if pc.Port != "5899" {
		t.Fatalf("Port mismatch: Expected %s, got %s\n", "5899", pc.Port)
	}
}

func TestParseSessionConfig(t *testing.T) {

	name := "test_s.json"
	f, err := os.Create(name)
	if err != nil {
		t.Fatalf(err.Error())
	}

	d := `{
	"Account": "alice",
	"Remote": "192.13.14.21",
	"KeyHome": "/home/sapta/vineyard",
	"Port": "5989"
}`

	f.Write([]byte(d))

	pc, err := ParseSessionConfig(name)
	if err != nil {
		t.Fatalf(err.Error())
	}

	os.Remove(name)

	if pc.Account != "alice" {
		t.Fatalf("Account mismatch: Expected %s, got %s\n", "alice", pc.Account)
	}

	if pc.Remote != "192.13.14.21" {
		t.Fatalf("Remote mismatch: Expected %s, got %s\n", "192.13.14.21", pc.Remote)
	}

	if pc.KeyHome != "/home/sapta/vineyard" {
		t.Fatalf("KeyHome mismatch: Expected %s, got %s\n", "/home/sapta/vineyard", pc.KeyHome)
	}

	if pc.Port != "5989" {
		t.Fatalf("Port mismatch: Expected %s, got %s\n", "5989", pc.Port)
	}

}
