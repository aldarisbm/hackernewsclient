package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HackerNewsClient struct {
	baseUrl string
	client  *http.Client
}

func (c *HackerNewsClient) GetTopStoriesIDs() ([]int, error) {
	topStoriesEndpoint := fmt.Sprintf("%s/%s", c.baseUrl, "topstories.json")

	res, err := c.client.Get(topStoriesEndpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var stories []int
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&stories); err != nil {
		return nil, err
	}
	return stories, nil
}

func (c *HackerNewsClient) GetItem(id int) (*Item, error) {
	getItemEndpoint := fmt.Sprintf("%s/%s/%d.json", c.baseUrl, "item", id)
	res, err := c.client.Get(getItemEndpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var item Item
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&item); err != nil {
		return nil, err
	}
	return &item, nil
}

func New(baseUrl string) *HackerNewsClient {
	return &HackerNewsClient{
		baseUrl: baseUrl,
		client:  http.DefaultClient,
	}
}
