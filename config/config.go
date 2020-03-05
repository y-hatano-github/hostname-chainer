package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HostName        string `envconfig:"HOSTNAME" required:"true"`
	NextHostAddress string `envconfig:"NEXT_HOST_ADDRESS" default:""`
	NextPort        string `envconfig:"NEXT_PORT" default:"8080"`
}

func (c *Config) LoadConfig() error {

	if err := envconfig.Process("", c); err != nil {
		return err
	}
	return nil
}
