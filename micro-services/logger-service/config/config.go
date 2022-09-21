package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InternalPort string `yaml:"internalPort"`
	MongoUser    string `yaml:"mongoUser"`
	MongoPass    string `yaml:"mongoPass"`
	LocalDbUrl   string `yaml:"localDbUrl"`
	DbUrl        string `yaml:"dbUrl"`
	RpcPort      string `yaml:"rpcPort"`
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
