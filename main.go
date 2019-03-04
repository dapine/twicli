package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	rootURL  = "https://api.twitch.tv/helix"
	clientId = "bhm2qd7pfnpy0l60y6rzlol6itbby7"
)

type Client struct {
	RootURL    string
	ClientId   string
	HttpClient *http.Client
}

type Game struct {
	Id        string `json: "id"`
	Name      string `json: "name"`
	BoxArtUrl string `json: "box_art_url"`
}

type Pagination struct {
	Cursor string `json: "cursor"`
}

type Games struct {
	Data       []Game     `json: "data"`
	Pagination Pagination `json : "pagination"`
}

func main() {
	client := NewClient(clientId)

	games, err := client.GetTopGames()
	if err != nil {
		log.Fatal(err)
	}

	for _, g := range games.Data {
		fmt.Println(g.Name)
	}
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

func (c *Client) GetTopGames() (*Games, error) {
	games := new(Games)

	res, err := c.Get("/games/top")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	s, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(s, games)

	return games, nil
}
