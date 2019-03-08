package main

import "net/http"

type Client struct {
	RootURL    string
	ClientId   string
	HttpClient *http.Client
}

func NewClient(clientId string) Client {
	// TODO: add options for http client
	httpc := &http.Client{}

	return Client{
		RootURL:    rootURL,
		ClientId:   clientId,
		HttpClient: httpc,
	}
}

func (c *Client) Get(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.RootURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Client-Id", c.ClientId)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
