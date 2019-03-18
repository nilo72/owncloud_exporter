package owncloud

import (
	"net/http"
	"strings"
)

// Client is a client for the ownCloud API.
type Client struct {
	httpClient *http.Client
	endpoint   string
	username   string
	password   string

	User UserClient
}

// A ClientOption is used to configure a Client.
type ClientOption func(*Client)

// WithEndpoint configures a Client to use the specified API endpoint.
func WithEndpoint(endpoint string) ClientOption {
	return func(client *Client) {
		client.endpoint = strings.TrimRight(endpoint, "/")
	}
}

// WithUsername configures a Client to use the specified username for authentication.
func WithUsername(username string) ClientOption {
	return func(client *Client) {
		client.username = username
	}
}

// WithPassword configures a Client to use the specified password for authentication.
func WithPassword(password string) ClientOption {
	return func(client *Client) {
		client.password = password
	}
}

// NewClient creates a new client.
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: &http.Client{},
	}

	for _, option := range options {
		option(client)
	}

	client.User = UserClient{client: client}

	return client
}
