package main

import (
	"strconv"
)

type Stream struct {
	Id           string        `json:"id"`
	UserId       string        `json:"user_id"`
	UserName     string        `json:"user_name"`
	GameId       string        `json:"game_id"`
	CommunityIds []interface{} `json:"community_ids"`
	Type         string        `json:"type"`
	Title        string        `json:"title"`
	ViewerCount  float32       `json:"viewer_count"`
	StartedAt    string        `json:"started_at"`
	Language     string        `json:"language"`
	ThumbnailUrl string        `json:"thumbnail_url"`
	TagIds       []interface{} `json:"tag_ids"`
}

type Streams struct {
	Data       []Stream   `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// Most viewed streams. n limit = 100
func (c *Client) GetTopStreams(n uint64) (*Streams, error) {
	// TODO: support more params
	streams := new(Streams)

	res, err := c.Get("/streams?first=" + strconv.FormatUint(n, 10))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = FromResponseJSON(res, streams)
	if err != nil {
		return nil, err
	}

	return streams, nil
}
