package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HostName        string `envconfig:"HOSTNAME" required:"true"`
	ListenPort      string `envconfig:"LISTEN_PORT" default:"8080"`
	NextHostAddress string `envconfig:"NEXT_HOST_ADDRESS" default:""`
	NextHostPort    string `envconfig:"NEXT_HOST_PORT" default:"8080"`
}

func (c *Config) LoadConfig() error {

	if err := envconfig.Process("", c); err != nil {
		return err
	}
	return nil
}
