package coc

import (
	"encoding/json"
	"strings"

	"github.com/clashgolang/coc/pkg/config"
	"github.com/clashgolang/coc/pkg/rest"
	log "github.com/sirupsen/logrus"
)

// Location is information about a location
type Location struct {
	CountryCode   string `json:"countryCode"`
	ID            int    `json:"id"`
	IsCountry     bool   `json:"isCountry"`
	LocalizedName string `json:"localizedName"`
	Name          string `json:"name"`
}

// String returns a string representation of a location
func (l Location) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// GetLocation gets the location information
func GetLocation(id string) (*Location, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations/")
	sb.WriteString(fmtTag(id))
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Location Location `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	// Return the clan
	return &resp.Location, nil
}

// GetLocations lists locations
func GetLocations(qparms rest.QParms) ([]Location, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations")

	body, err := get(sb.String(), nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Locations []Location `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Locations, nil
}
