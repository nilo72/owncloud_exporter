package client

import "time"

type ClientConfig struct {
	Timeout time.Duration
	Address string
}

type Client struct {

}

func NewClient(cfg *ClientConfig) *Client {
	return &Client{}
}
