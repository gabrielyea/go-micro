package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	LocalURL  string `yaml:"localurl"`
	RemoteURL string `yaml:"remoteurl"`
}

func NewConfig(fileName string) (*Config, error) {
	c := &Config{}
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
