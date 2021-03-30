package coc

import (
	"encoding/json"
	"strings"

	"github.com/clashgolang/coc/pkg/config"
	"github.com/clashgolang/coc/pkg/rest"
	log "github.com/sirupsen/logrus"
)

// League lists leagues
type League struct {
	IconUrls IconUrls `json:"iconUrls"`
	ID       int      `json:"id"`
	Name     string   `json:"name"`
}

// String returns a string representation of a league
func (l League) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// LeagueSeason is a league season.
type LeagueSeason struct {
	ID string `json:"id"`
}

// String returns a string representation of a league season
func (ls LeagueSeason) String() string {
	b, _ := json.Marshal(ls)
	return string(b)
}

// LeagueSeasonRanking is the league season ranking.
type LeagueSeasonRanking struct {
	AttackWins   int           `json:"attackWins"`
	Clan         ClanReference `json:"clan"`
	DefenseWins  int           `json:"defenseWins"`
	ExpLevel     int           `json:"expLevel"`
	League       League        `json:"league"`
	Name         string        `json:"name"`
	PreviousRank int           `json:"previousRank"`
	Rank         int           `json:"rank"`
	Tag          string        `json:"tag"`
	Trophies     int           `json:"trophies"`
}

// String returns a string representation of a league season
func (lsr LeagueSeasonRanking) String() string {
	b, _ := json.Marshal(lsr)
	return string(b)
}

// WarLeague is information about a war league.
type WarLeague struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// String returns a string representation of a war league
func (wl WarLeague) String() string {
	b, _ := json.Marshal(wl)
	return string(b)
}

// GetLeague gets the league information
func GetLeague(leagueID string) (*League, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/leagues/")
	sb.WriteString(fmtTag(leagueID))
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		League League `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return &resp.League, nil
}

// GetLeagues lists the leagues
func GetLeagues(qparms rest.QParms) ([]League, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/leagues/")

	body, err := get(sb.String(), nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Leagues []League `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Leagues, nil
}

// GetLeagueSeasons gets the league seasons
func GetLeagueSeasons(leagueID string) ([]LeagueSeason, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/leagues/")
	sb.WriteString(fmtTag(leagueID))
	sb.WriteString("/seasons")
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Seasons []LeagueSeason `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Seasons, nil
}

// GetLeagueSeasonRankings gets the league season rankings for Legend League
func GetLeagueSeasonRankings(leagueID string) ([]LeagueSeasonRanking, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/leagues/")
	sb.WriteString(fmtTag(leagueID))
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Rankings []LeagueSeasonRanking `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Rankings, nil
}

// GetLWareague gets the war league information
func GetWarLeague(leagueID string) (*WarLeague, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/warleagues/")
	sb.WriteString(fmtTag(leagueID))
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		League WarLeague `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return &resp.League, nil
}

// GetWarLeagues lists the war leagues
func GetWarLeagues(qparms rest.QParms) ([]WarLeague, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/warleagues/")
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Leagues []WarLeague `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	return resp.Leagues, nil
}
