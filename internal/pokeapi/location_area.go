package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationAreaName
	//fmt.Println(url)

	var data []byte
	var ok bool
	data, ok = c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		c.cache.Add(url, data)
	} else {
		fmt.Printf("url found in cache: %s\n", url)
	}
	location := LocationArea{}

	err := json.Unmarshal(data, &location)
	if err != nil {
		return LocationArea{}, err
	}

	return location, nil
}
