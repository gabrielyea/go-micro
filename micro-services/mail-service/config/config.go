package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	WebPort     string `yaml:"webPort"`
	MailDomain  string `yaml:"mailDomain"`
	MailHost    string `yaml:"mailHost"`
	MailPort    string `yaml:"mailPort"`
	UserName    string `yaml:"userName"`
	Password    string `yaml:"password"`
	Encryption  string `yaml:"encryption"`
	FromName    string `yaml:"fromName"`
	FromAddress string `yaml:"fromAddress"`
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
