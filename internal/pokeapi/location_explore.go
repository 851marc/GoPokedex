package pokeapi

import (
	"encoding/json"
)

// ListLocations -
func (c *Client) ExploreLocation(region string) (Location, error) {
	url := baseURL + "/location-area/" + region

	dat, err := c.Get(url)
	if err != nil {
		return Location{}, err
	}

	locationsResp := Location{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	return locationsResp, nil
}
