package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/TomasMoralesT/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	baseURL    string
	cache      *pokecache.Cache
}

// NewClient creates a new PokeAPI client
func NewClient() Client {
	return Client{
		httpClient: http.Client{},
		baseURL:    "https://pokeapi.co/api/v2",
		cache:      pokecache.NewCache(5 * time.Minute),
	}
}

func (c *Client) GetLocationArea(pageURL *string) (LocationAreaList, error) {
	endpoint := "/location-area"
	fullURL := c.baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	if cached, found := c.cache.Get(fullURL); found {
		var result LocationAreaList
		err := json.Unmarshal(cached, &result)
		return result, err
	}

	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		return LocationAreaList{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaList{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	c.cache.Add(fullURL, body)

	var result LocationAreaList
	if err := json.Unmarshal(body, &result); err != nil {
		return LocationAreaList{}, err
	}
	return result, nil
}

func (c *Client) GetLocationAreaByName(name string) (LocationArea, error) {
	endpoint := fmt.Sprintf("/location-area/%s", name)
	fullURL := c.baseURL + endpoint

	if cached, found := c.cache.Get(fullURL); found {
		var result LocationArea
		err := json.Unmarshal(cached, &result)
		return result, err
	}

	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, body)

	var result LocationArea
	if err := json.Unmarshal(body, &result); err != nil {
		return LocationArea{}, err
	}

	return result, nil

}

func (c *Client) GetPokemonData(name string) (Pokemon, error) {
	endpoint := fmt.Sprintf("/pokemon/%s", name)
	fullURL := c.baseURL + endpoint

	if cached, found := c.cache.Get(fullURL); found {
		var result Pokemon
		err := json.Unmarshal(cached, &result)
		return result, err
	}

	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, body)

	var result Pokemon
	if err := json.Unmarshal(body, &result); err != nil {
		return Pokemon{}, err
	}

	return result, nil

}
