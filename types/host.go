package types

import (
	"hostname-chainer/config"
)

type CurrentHost interface {
	GetHostName() string
	GetNextHostAddress() string
	GetNextHostPort() string
	HasNextHostInfo() bool
}

type currentHostInfo struct {
	hostName        string
	nextHostPort    string
	nextHostAddress string
}

func NewHost(conf config.Config) CurrentHost {
	return &currentHostInfo{
		hostName:        conf.HostName,
		nextHostAddress: conf.NextHostAddress,
		nextHostPort:    conf.NextHostPort,
	}
}

func (h *currentHostInfo) GetHostName() string {
	return h.hostName
}

func (h *currentHostInfo) GetNextHostAddress() string {
	return h.nextHostAddress
}

func (h *currentHostInfo) GetNextHostPort() string {
	return h.nextHostPort
}

func (h *currentHostInfo) HasNextHostInfo() bool {
	return h.nextHostAddress != "" && h.nextHostPort != ""
}
