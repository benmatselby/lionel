package trello

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

const baseURL = "https://api.trello.com/1"

// API defines the interface for the client
type API interface {
	GetBoard(name string) (*Board, error)
	GetBoards() ([]Board, error)
	GetCards(board Board) ([]Card, error)
	GetLists(board Board) ([]List, error)
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

// Get is responsible for performing a GET request
func (c *Client) Get(url string, response interface{}) (*http.Response, error) {
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

	url := "/members/me/boards"
	_, err := c.Get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("Unable to fulfil request %s: %s", url, err)
	}

	return response, nil
}

// GetBoard will return a single board
func (c *Client) GetBoard(name string) (*Board, error) {
	boards, err := c.GetBoards()
	if err != nil {
		return nil, err
	}

	for _, board := range boards {
		if board.Name == name {
			return &board, nil
		}
	}

	return nil, fmt.Errorf("Unable to find board with name '%s'", name)
}

// Card defines what a single card looks like
type Card struct {
	Name   string `json:"name"`
	ListID string `json:"idList"`
}

// GetCards will return a set of cards for a given board name
func (c *Client) GetCards(board Board) ([]Card, error) {
	var response []Card

	url := fmt.Sprintf("/boards/%s/cards", board.ID)
	_, err := c.Get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("Unable to fulfil request %s: %s", url, err)
	}

	return response, nil
}

// List defines a column on the board
type List struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetLists will return all the columns for a given board
func (c *Client) GetLists(board Board) ([]List, error) {
	var response []List

	url := fmt.Sprintf("/boards/%s/lists", board.ID)
	_, err := c.Get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("Unable to fulfil request %s: %s", url, err)
	}

	return response, nil
}
