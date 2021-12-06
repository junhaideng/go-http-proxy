package config

import (
	"os"

	"github.com/junhaideng/go-proxy/http"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server *Server `json:"server" yaml:"server"`
}

type Server struct {
	Port       uint16     `json:"port" yaml:"port"`
	Host       string     `json:"host" yaml:"host"`
	Auth       *http.Auth `json:"auth" yaml:"auth"`
	EnableAuth bool       `json:"enable_auth" yaml:"enable_auth"`
}

func New(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config

	yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, nil
}
