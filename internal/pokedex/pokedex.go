package pokedex

import (
	"encoding/json"
	"errors"
	"net/http"
)

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type NamedAPIResourceList struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

func Map(next string) (NamedAPIResourceList, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if next != "" {
		url = next
	}
	res, err := http.Get(url)
	if err != nil {
		return NamedAPIResourceList{}, err
	}
	defer res.Body.Close()

	var locations NamedAPIResourceList
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return NamedAPIResourceList{}, err
	}
	return locations, nil
}

func Mapb(previous string) (NamedAPIResourceList, error) {
	url := previous
	if url == "" {
		return NamedAPIResourceList{}, errors.New("previous shouldn't be empty, you are on the first page")
	}
	res, err := http.Get(url)
	if err != nil {
		return NamedAPIResourceList{}, err
	}
	defer res.Body.Close()

	var locations NamedAPIResourceList
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	return locations, nil
}
