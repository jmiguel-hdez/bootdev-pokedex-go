package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (NamedAPIResourceList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return NamedAPIResourceList{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	locationsResp := NamedAPIResourceList{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return NamedAPIResourceList{}, err
	}
	return locationsResp, nil
}
