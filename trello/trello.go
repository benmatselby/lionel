package trello

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

const baseURL = "https://api.trello.com/1/"

// API defines the interface for the client
type API interface {
	GetBoards() ([]Board, error)
}

// Client is the Trello concrete implementation
type Client struct {
	Client  *http.Client
	BaseURL string
	Key     string
	Token   string
	ctx     context.Context
}

// New will return you a Trello client
func New() Client {
	client := Client{
		Client: http.DefaultClient,
		Key:    viper.GetString("TRELLO_CLI_KEY"),
		Token:  viper.GetString("TRELLO_CLI_SECRET"),
	}
	return client
}

func (c *Client) get(url string, response interface{}) (*http.Response, error) {
	request, err := http.NewRequest("GET", baseURL+url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "benmatselby/lionel")

	q := request.URL.Query()
	q.Add("key", c.Key)
	q.Add("token", c.Token)
	request.URL.RawQuery = q.Encode()

	httpRes, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode != 200 {
		return nil, fmt.Errorf("Request to %s responded with status %d", request.URL, httpRes.StatusCode)
	}

	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("Decoding json response from %s failed: %v", request.URL, err)
	}

	return httpRes, nil
}

// Board defines what a single board looks like
type Board struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Closed bool   `json:"closed"`
}

// GetBoards will return a list boards the user can access
func (c *Client) GetBoards() ([]Board, error) {
	var response []Board

	url := "members/me/boards"
	_, err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("Unable to fulfil request %s: %s", url, err)
	}

	return response, nil
}
