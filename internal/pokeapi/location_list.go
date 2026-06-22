package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (NamedAPIResourceList, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return NamedAPIResourceList{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return NamedAPIResourceList{}, err
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return NamedAPIResourceList{}, err
		}
		c.cache.Add(url, data)
	} else {
		fmt.Printf("getting data from cache: key %s\n", url)
	}

	locationsResp := NamedAPIResourceList{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return NamedAPIResourceList{}, err
	}
	return locationsResp, nil
}
