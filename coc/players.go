package coc

import (
	"encoding/json"
	"strings"

	"github.com/clashgolang/coc/pkg/config"
	"github.com/clashgolang/coc/pkg/rest"
	log "github.com/sirupsen/logrus"
)

// Player is a single player in Clash of Clans.
type Player struct {
	Achievements         []PlayerAchievement `json:"achievements"`
	AttackWins           int                 `json:"attackWins"`
	BestTrophies         int                 `json:"bestTrophies"`
	BestVersusTrophies   int                 `json:"bestVersusTrophies"`
	BuilderHallLevel     int                 `json:"builderHallLevel"`
	Clan                 ClanReference       `json:"clan"`
	DefenseWins          int                 `json:"defenseWins"`
	Donations            int                 `json:"donations"`
	DonationsReceived    int                 `json:"donationsReceived"`
	ExpLevel             int                 `json:"expLevel"`
	Heroes               []Troop             `json:"heroes"`
	Labels               []Label             `json:"labels"`
	League               League              `json:"league"`
	Name                 string              `json:"name"`
	Role                 string              `json:"role"`
	Spells               []Troop             `json:"spells"`
	Tag                  string              `json:"tag"`
	TownHallLevel        int                 `json:"townHallLevel"`
	Troops               []Troop             `json:"troops"`
	Trophies             int                 `json:"trophies"`
	VersusBattleWinCount int                 `json:"versusBattleWinCount"`
	VersusBattleWins     int                 `json:"versusBattleWins"`
	VersusTrophies       int                 `json:"versusTrophies"`
	WarStars             int                 `json:"warStars"`
}

// String returns a string representation of a player
func (p Player) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

type PlayerAchievement struct {
	CompletionInfo string `json:"completionInfo"`
	Info           string `json:"info"`
	Name           string `json:"name"`
	Stars          int    `json:"stars"`
	Target         int    `json:"target"`
	Value          int    `json:"value"`
	Village        string `json:"village"`
}

// String returns a string representation of a player achievement
func (a PlayerAchievement) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}

// Troop represents a troop, hero or spell in Clash of Clans
type Troop struct {
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Name     string `json:"name"`
	Village  string `json:"village"`
}

// String returns a string representation of a troop
func (t Troop) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

// PlayerRanking is the ranking of a player for specific location.
type PlayerRanking struct {
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
func (l PlayerRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// PlayerVersusRanking is the player ranking for a specific location
type PlayerVersusRanking struct {
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
func (l PlayerVersusRanking) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// GetPlayer retrieves information about a given player
func GetPlayer(tag string) (*Player, error) {
	// Build the URL
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/players/")
	sb.WriteString(fmtTag(tag))
	url := sb.String()
	log.Trace(url)

	// Get the player
	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}
	var player Player
	if err := json.Unmarshal(body, &player); err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	// Return the player
	return &player, nil
}

// GetPlayerRankings gets player rankings for a specific location
func GetPlayerRankings(locationID string, qparms rest.QParms) ([]PlayerRanking, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/locations/")
	sb.WriteString(fmtTag(locationID))
	sb.WriteString("/rankings/players")
	url := sb.String()
	log.Trace(url)

	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		Rankings []PlayerRanking `json:"items"`
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

// GetPlayerRankings gets clan versus rankings for a specific location
func GetPlayerVersusRankings(locationID string, qparms rest.QParms) ([]PlayerVersusRanking, error) {
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
		Rankings []PlayerVersusRanking `json:"items"`
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
