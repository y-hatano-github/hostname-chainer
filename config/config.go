package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HostName        string `envconfig:"HOSTNAME" required:"true"`
	DestinationHost string `envconfig:"DESTINATION_HOST" default:""`
	DestinationPort string `envconfig:"DESTINATION_PORT" default:""`
}

func (c *Config) LoadConfig() error {

	if err := envconfig.Process("", c); err != nil {
		return err
	}
	return nil
}
