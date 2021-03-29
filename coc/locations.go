package coc

import (
	"encoding/json"
	"net/url"
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

// LocationClanRanking is the clan ranking for a specific location.
type LocationClanRanking struct {
	BadgeUrls    BadgeUrls `json:"badgeUrls"`
	ClanLevel    int       `json:"clanLevel"`
	ClanPoints   int       `json:"clanPoints"`
	Location     Location  `json:"location"`
	Members      int       `json:"members"`
	Name         string    `json:"name"`
	PreviousRank int       `json:"previousRank"`
	Rank         int       `json:"rank"`
	Tag          string    `json:"tag"`
}

// String returns a string representation of a location clan ranking
func (l LocationClanRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LocationPlayerRanking is the ranking of a player for specific location.
type LocationPlayerRanking struct {
	Clan         ClanReference `json:"clan"`
	League       League        `json:"league"`
	AttackWins   int           `json:"attackWins"`
	DefenseWins  int           `json:"defenseWins"`
	Tag          string        `json:"tag"`
	Name         string        `json:"name"`
	ExpLevel     int           `json:"expLevel"`
	Rank         int           `json:"rank"`
	PreviousRank int           `json:"previousRank"`
	Trophies     int           `json:"trophies"`
}

// String returns a string representation of a location player ranking
func (l LocationPlayerRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LocationClanVersusRanking is the clan versus ranking for a specific location
type LocationClanVersusRanking struct {
	ClanVersusPoints int `json:"clanVersusPoints"`
	ClanPoints       int `json:"clanPoints"`
}

// String returns a string representation of a clan-versus ranking for a location
func (l LocationClanVersusRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LocationPlayerVersusRanking is the player ranking for a specific location
type LocationPlayerVersusRanking struct {
	Clan             ClanReference `json:"clan"`
	VersusBattleWins int           `json:"versusBattleWins"`
	Tag              string        `json:"tag"`
	Name             string        `json:"name"`
	ExpLevel         int           `json:"expLevel"`
	Rank             int           `json:"rank"`
	PreviousRank     int           `json:"previousRank"`
	VersusTrophies   int           `json:"versusTrophies"`
}

// String returns a string representation of a player-versus ranking for a location
func (l LocationPlayerVersusRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

func GetLocation(id string) (*Location, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations/")

	// If the tag doesn't have a '#' character at the front, add one
	if id[0] == '#' {
		sb.WriteString(url.QueryEscape(id))
	} else {
		sb.WriteString(url.QueryEscape("#" + id))
	}

	body, err := get(sb.String(), nil)
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
