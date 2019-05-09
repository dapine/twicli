package main

type Game struct {
	Id        string `json: "id"`
	Name      string `json: "name"`
	BoxArtUrl string `json: "box_art_url"`
}

type Games struct {
	Data       []Game     `json: "data"`
	Pagination Pagination `json : "pagination"`
}

func (c *Client) GetGames(endpoint string) (*Games, error) {
	games := new(Games)

	res, err := c.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = FromResponseJSON(res, games)
	if err != nil {
		return nil, err
	}

	return games, nil

}

func (c *Client) GetTopGames() (*Games, error) {
	games, err := c.GetGames("/games/top")
	if err != nil {
		return nil, err
	}

	return games, nil
}

func (c *Client) Paginate(url, before, after string) (*Games, error) {
	games, err := c.GetGames(url + "?before=" + before + "&after=" + after)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func (c *Client) GetGamesFromId(id string) (*Games, error) {
	games, err := c.GetGames("/games?id=" + id)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func (c *Client) GetGamesFromName(name string) (*Games, error) {
	games, err := c.GetGames("/games?name=" + name)
	if err != nil {
		return nil, err
	}

	return games, nil
}
