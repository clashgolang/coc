package coc

import (
	"encoding/json"
	"strings"

	"github.com/clashgolang/coc/pkg/config"
	"github.com/clashgolang/coc/pkg/rest"
	log "github.com/sirupsen/logrus"
)

// Label is a label for a clan or player.
type Label struct {
	Name     string   `json:"name"`
	ID       int      `json:"id"`
	IconUrls IconUrls `json:"iconUrls"`
}

// String returns a string representation of a label
func (l Label) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// GetClanLabels lists clan labels
func GetClanLabels(qparms rest.QParms) ([]Label, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/labels/clans/")

	body, err := get(sb.String(), nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Labels []Label `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Labels, nil
}

// GetPlayerLabels lists player labels
func GetPlayerLabels(qparms rest.QParms) ([]Label, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/labels/players/")

	body, err := get(sb.String(), nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Labels []Label `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Labels, nil
}
