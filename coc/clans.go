package coc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/clashgolang/coc/pkg/config"
	"github.com/clashgolang/coc/pkg/rest"

	log "github.com/sirupsen/logrus"
)

// Clan is a clan in Clash of Clans.
type Clan struct {
	BadgeUrls        BadgeUrls     `json:"badgeUrls"`
	ClanLevel        int           `json:"clanLevel"`
	ClanPoints       int           `json:"clanPoints"`
	ClanVersusPoints int           `json:"clanVersusPoints"`
	Description      string        `json:"description"`
	IsWarLogPublic   bool          `json:"isWarLogPublic"`
	Labels           []Label       `json:"labels"`
	Location         Location      `json:"location"`
	Members          int           `json:"members"`
	Name             string        `json:"name"`
	RequiredTrophies int           `json:"requiredTrophies"`
	Tag              string        `json:"tag"`
	Type             string        `json:"type"`
	WarFrequency     string        `json:"warFrequency"`
	WarLeague        ClanWarLeague `json:"warLeague"`
	WarLosses        int           `json:"warLosses"`
	WarTies          int           `json:"warTies"`
	WarWins          int           `json:"warWins"`
	WarWinStreak     int           `json:"warWinStreak"`
}

// String returns a string representation of a clan
func (c Clan) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// ClanMember is a member of a given clan.
type ClanMember struct {
	ClanRank          int    `json:"clanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
	ExpLevel          int    `json:"expLevel"`
	League            League `json:"league"`
	Name              string `json:"name"`
	PreviousClanRank  int    `json:"previousClanRank"`
	Role              string `json:"role"`
	Tag               string `json:"tag"`
	Trophies          int    `json:"trophies"`
	VersusTrophies    int    `json:"versusTrophies"`
}

// String returns a string representation of a clan member
func (m ClanMember) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

// ClanRanking is the clan ranking for a specific location.
type ClanRanking struct {
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

// String returns a string representation of a clan ranking
func (l ClanRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// ClanVersusRanking is the clan versus ranking for a specific location
type ClanVersusRanking struct {
	ClanVersusPoints int `json:"clanVersusPoints"`
	ClanPoints       int `json:"clanPoints"`
}

// String returns a string representation of a clan versus ranking
func (l ClanVersusRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// ClanReference provides a reference to a given clan
type ClanReference struct {
	BadgeUrls BadgeUrls `json:"badgeUrls"`
	ClanLevel int       `json:"clanLevel"`
	Name      string    `json:"name"`
	Tag       string    `json:"tag"`
}

// String returns a string representation of a clan member
func (r ClanReference) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// GetClan retrieves information about a clan with the given tag
func GetClan(tag string) (*Clan, error) {
	// Build the URL
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(tag))
	url := sb.String()
	log.Trace(url)

	// Get the clan
	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}
	var clan Clan
	if err := json.Unmarshal(body, &clan); err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	// Return the clan
	return &clan, nil
}

// GetClans returns information about all clans that match the query parameters
func GetClans(name string, qparms rest.QParms) ([]Clan, error) {
	url := config.Data.BaseURL + "/clans"
	body, err := get(url, qparms)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Clans []Clan `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Clans, nil
}

// GetClanMembers gets information about members of a given clan
func GetClanMembers(clanTag string, qparms rest.QParms) ([]ClanMember, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(clanTag))
	sb.WriteString("/members")

	body, err := get(sb.String(), qparms)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	// Parse into an array of clans
	type respType struct {
		ClanMembers []ClanMember `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.ClanMembers, nil
}

// GetClan Rankings gets clan rankings for a specific location
func GetClanRankings(locationID string, qparms rest.QParms) ([]ClanRanking, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations/")
	sb.WriteString(fmtTag(locationID))
	sb.WriteString("/rankings/clans")
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clan rankings
	type respType struct {
		Rankings []ClanRanking `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	// Return the clan
	return resp.Rankings, nil
}

// GetClanRankings gets clan versus rankings for a specific location
func GetClanVersusRankings(locationID string, qparms rest.QParms) ([]ClanVersusRanking, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations/")
	sb.WriteString(fmtTag(locationID))
	sb.WriteString("/rankings/clan-versus")
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clan rankings
	type respType struct {
		Rankings []ClanVersusRanking `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	// Return the clan
	return resp.Rankings, nil
}
