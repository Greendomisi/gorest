package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerAdr  string `json:"ServerAdr"`
	PostgreAdr string `json:"PostgreAdr"`
}

func New(path string) (*Config, error) {
	c := &Config{}

	f, err := os.Open(path)
	if err != nil {
		return &Config{}, err
	}

	defer f.Close()

	d := json.NewDecoder(f)

	err = d.Decode(c)
	if err != nil {
		return &Config{}, err
	}

	return c, nil
}

// func (c *Config) parse() (err error) {

// }
