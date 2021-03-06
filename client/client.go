package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// next host interface
type Client interface {
	Post(hosts *Hosts) (*Hosts, error)
}

type client struct {
	baseUrl string
	client  *resty.Client
}

// create next host client
func NewClient(nextAddress, nextPort string) Client {
	return &client{
		baseUrl: fmt.Sprintf("http://%s:%s", nextAddress, nextPort),
		client:  resty.New(),
	}
}

func (h *client) Post(hosts *Hosts) (*Hosts, error) {
	resp, err := h.client.R().SetBody(hosts).Post(h.baseUrl)
	if err != nil {
		return nil, err
	}

	var response Hosts
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

//  host info
type Host struct {
	HostName string `json:"host_name"`
	To       string `json:"to,omitempty"`
}

// type host array
type Hosts []Host

// add current host info to Hosts
func (r *Hosts) AddCurrentHostInfo(hostName string, nextHost string) *Hosts {
	*r = append(*r, Host{HostName: hostName, To: nextHost})
	return r
}
