package nexthost

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client interface {
	RequestToNextHost() (*resty.Response, error)
}

type client struct {
	baseUrl string
	client  *resty.Client
}

func NewNextHostClient(nextAddress, nextPort string) Client {

	return &client{
		baseUrl: fmt.Sprintf("http://%s:%s", nextAddress, nextPort),
		client:  resty.New(),
	}
}

func (h *client) RequestToNextHost() (*resty.Response, error) {
	return h.client.R().Get(h.baseUrl)
}
