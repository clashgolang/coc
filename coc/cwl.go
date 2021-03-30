package coc

import (
	"encoding/json"
	"strings"

	"github.com/clashgolang/coc/pkg/config"
	log "github.com/sirupsen/logrus"
)

// ClanWarLeague is a reference to a given clan war league
type ClanWarLeague struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// String returns a string representation of a clan war league
func (wl ClanWarLeague) String() string {
	b, _ := json.Marshal(wl)
	return string(b)
}

// ClanWarLeagueGroup is a clan's current clan war league group.
type ClanWarLeagueGroup struct {
	Clans []struct {
		BadgeUrls BadgeUrls `json:"badgeUrls"`
		ClanLevel int       `json:"clanLevel"`
		Members   []struct {
			Name          string `json:"name"`
			Tag           string `json:"tag"`
			TownHallLevel int    `json:"townHallLevel"`
		} `json:"members"`
		Name string `json:"name"`
		Tag  string `json:"tag"`
	} `json:"clans"`
	Rounds []struct {
		WarTags []string `json:"warTags"`
	} `json:"rounds"`
	Season string `json:"season"`
	State  string `json:"state"`
	Tag    string `json:"tag"`
}

// String returns a string representation of a clan war league group
func (lg ClanWarLeagueGroup) String() string {
	b, _ := json.Marshal(lg)
	return string(b)
}

// ClanWarLeagueWar is information about an individual clan war league war
type ClanWarLeagueWar struct {
	Clan                 ClanWarTeam `json:"clan"`
	EndTime              string      `json:"endTime"`
	Opponent             ClanWarTeam `json:"opponent"`
	PreparationStartTime string      `json:"preparationStartTime"`
	StartTime            string      `json:"startTime"`
	State                string      `json:"state"`
	TeamSize             int         `json:"teamSize"`
	WarStartTime         string      `json:"warStartTime"`
}

// String returns a string representation of a clan war league war
func (lw ClanWarLeagueWar) String() string {
	b, _ := json.Marshal(lw)
	return string(b)
}

// GetClanWarLeagueGroup retrieves information about clan's current clan war league group
func GetClanWarLeagueGroup(clanTag string) (*ClanWarLeagueGroup, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(clanTag))
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		LeagueGroup ClanWarLeagueGroup `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return &resp.LeagueGroup, nil
}

// GetClanWarLeagueWar retrieves information about individual clan war league war
func GetClanWarLeagueWar(clanTag string) (*ClanWarLeagueWar, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clanswarleagues/wars/")
	sb.WriteString(fmtTag(clanTag))
	sb.WriteString("/currentwar")
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		LeagueWar ClanWarLeagueWar `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return &resp.LeagueWar, nil
}
