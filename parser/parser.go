package parser

import (
	"encoding/json"

	"github.com/sap200/dvpn-node/utils"
)

// ServerConfig is a struct that contains configuration of server
type ServerConfig struct {
	Type    string `json:"Type"`
	Account string `json:"Account"`
	Remote  string `json:"Remote"`
	KeyHome string `json:"KeyHome"`
	Port    string `json:"Port"`
}

// SessionConfig is a struct that contains configuration of a connection session
type SessionConfig struct {
	Type    string `json:"Type"`
	Account string `json:"Account"`
	Remote  string `json:"Remote"`
	KeyHome string `json:"KeyHome"`
	Port    string `json:"Port"`
	IPAddr  string `json:"IPAddr"`
}

// WebAppConfig contains configuration for web based interface
type WebAppConfig struct {
	Port    string `json:"Port"`
	QNode   string `json:"QNode"`
	Account string `json:"Account"`
	LogPath string `json:"LogPath"`
	KeyHome string `json:"KeyHome"`
}

// ParseServerConfig parses json file
func ParseServerConfig(path string) (*ServerConfig, error) {
	data, err := utils.ReadFile(path)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(data))

	// try unmarshalling
	var config ServerConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// ParseSessionConfig parses the connection config file
// It has a json structure
func ParseSessionConfig(path string) (*SessionConfig, error) {
	data, err := utils.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// try unmarshalling
	var config SessionConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// ParseWebAppConfig parses the web app configuration
func ParseWebAppConfig(path string) (*WebAppConfig, error) {
	data, err := utils.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// try unmarshalling
	var config WebAppConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
