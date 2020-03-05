package types

import (
	"recursive-echo/config"
	"recursive-echo/nexthost"
)

type Host interface {
	GetName() string
	HasNextHostInfo() bool
	GetNextHostClient() nexthost.Client
}

type HostInfo struct {
	hostName     string
	nextHostPort string
	nextHostAddress string
}

func NewHost(conf config.Config) Host {
	return &HostInfo{
		hostName:    conf.HostName,
		nextHostAddress: conf.NextHostAddress,
		nextHostPort: conf.NextPort,
	}
}

func (h *HostInfo) GetName() string {
	return h.hostName
}

func (h *HostInfo) HasNextHostInfo() bool {
	return h.nextHostAddress != "" && h.nextHostPort != ""
}

func (h *HostInfo) GetNextHostClient() nexthost.Client {
	return nexthost.NewNextHostClient(h.nextHostAddress, h.nextHostPort)
}
