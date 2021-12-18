package client

import (
	"net/http"
)

type Client struct {
	Host   string
	Token  string
	Client *http.Client
}

// NewClient returns a new sonnen reader client
func NewClient(host string, token string) *Client {
	return &Client{
		Host:   host,
		Token:  token,
		Client: &http.Client{},
	}
}

func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if c.Token != "" {
		req.Header.Set("Auth-Token", c.Token)
	}
	return c.Client.Do(req)
}
