package coc

import (
	"encoding/json"
	"strings"

	"github.com/clashgolang/coc/pkg/config"
	log "github.com/sirupsen/logrus"
)

// Player is a single player in Clash of Clans.
type Player struct {
	Achievements []struct {
		CompletionInfo string `json:"completionInfo"`
		Info           string `json:"info"`
		Name           string `json:"name"`
		Stars          int    `json:"stars"`
		Target         int    `json:"target"`
		Value          int    `json:"value"`
		Village        string `json:"village"`
	} `json:"achievements"`
	AttackWins         int `json:"attackWins"`
	BestTrophies       int `json:"bestTrophies"`
	BestVersusTrophies int `json:"bestVersusTrophies"`
	BuilderHallLevel   int `json:"builderHallLevel"`
	Clan               struct {
		BadgeUrls BadgeUrls `json:"badgeUrls"`
		ClanLevel int       `json:"clanLevel"`
		Name      string    `json:"name"`
		Tag       string    `json:"tag"`
	} `json:"clan"`
	DefenseWins          int     `json:"defenseWins"`
	Donations            int     `json:"donations"`
	DonationsReceived    int     `json:"donationsReceived"`
	ExpLevel             int     `json:"expLevel"`
	Heroes               []Troop `json:"heroes"`
	Labels               []Label `json:"labels"`
	League               League  `json:"league"`
	Name                 string  `json:"name"`
	Role                 string  `json:"role"`
	Spells               []Troop `json:"spells"`
	Tag                  string  `json:"tag"`
	TownHallLevel        int     `json:"townHallLevel"`
	Troops               []Troop `json:"troops"`
	Trophies             int     `json:"trophies"`
	VersusBattleWinCount int     `json:"versusBattleWinCount"`
	VersusBattleWins     int     `json:"versusBattleWins"`
	VersusTrophies       int     `json:"versusTrophies"`
	WarStars             int     `json:"warStars"`
}

// String returns a string representation of a player
func (p Player) String() string {
	b, _ := json.Marshal(p)
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
