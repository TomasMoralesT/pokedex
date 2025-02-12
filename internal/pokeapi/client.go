package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient http.Client
	baseURL    string
}

// NewClient creates a new PokeAPI client
func NewClient() Client {
	return Client{
		httpClient: http.Client{},
		baseURL:    "https://pokeapi.co/api/v2",
	}
}

func (c *Client) GetLocationArea(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := c.baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaResp{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	var result LocationAreaResp
	if err := json.Unmarshal(body, &result); err != nil {
		return LocationAreaResp{}, err
	}
	return result, nil
}
