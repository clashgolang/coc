package coc

import (
	"encoding/json"
	"strings"

	"github.com/clashgolang/coc/pkg/config"
	"github.com/clashgolang/coc/pkg/rest"
	log "github.com/sirupsen/logrus"
)

// ClanWar is a given war in a clan's war log.
type ClanWar struct {
	State                string      `json:"state,omitempty"`
	TeamSize             int         `json:"teamSize"`
	PreparationStartTime CoCTime     `json:"preparationStartTime,omitempty"`
	StartTime            CoCTime     `json:"startTime,omitempty"`
	EndTime              CoCTime     `json:"endTime,omitempty"`
	Result               string      `json:"result,omitempty"`
	Clan                 ClanWarTeam `json:"clan"`
	Opponent             ClanWarTeam `json:"opponent"`
}

// String returns a string representation of a clan war
func (cw ClanWar) String() string {
	b, _ := json.Marshal(cw)
	return string(b)
}

// ClanWarTeam is the clan that is participating in the clan war.
type ClanWarTeam struct {
	Attacks               int             `json:"attacks"`
	BadgeUrls             BadgeUrls       `json:"badgeUrls"`
	ClanLevel             int             `json:"clanLevel"`
	DestructionPercentage float32         `json:"destructionPercentage"`
	ExpEarned             int             `json:"expEarned"`
	Members               []ClanWarMember `json:"members,omitempty"`
	Name                  string          `json:"name"`
	Stars                 int             `json:"stars"`
	Tag                   string          `json:"tag"`
}

// String returns a string representation of a clan war team
func (cwt ClanWarTeam) String() string {
	b, _ := json.Marshal(cwt)
	return string(b)
}

// ClanWarMember is a member who participated in a clan war.
type ClanWarMember struct {
	Attacks            []ClanWarAttack `json:"attacks,omitempty"`
	BestOpponentAttack ClanWarAttack   `json:"bestOpponentAttack,omitempty"`
	MapPosition        int             `json:"mapPosition"`
	Name               string          `json:"name"`
	OpponentAttacks    int             `json:"opponentAttacks"`
	Tag                string          `json:"tag"`
	TownhallLevel      int             `json:"townhallLevel"`
}

// String returns a string representation of a clan war member
func (cwm ClanWarMember) String() string {
	b, _ := json.Marshal(cwm)
	return string(b)
}

// ClanWarAttack is an attack made in a clan war.
type ClanWarAttack struct {
	Order                 int    `json:"order"`
	AttackerTag           string `json:"attackerTag"`
	DefenderTag           string `json:"defenderTag"`
	Stars                 int    `json:"stars"`
	DestructionPercentage int    `json:"destructionPercentage"`
}

// String returns a string representation of a clan war atack
func (cwa ClanWarAttack) String() string {
	b, _ := json.Marshal(cwa)
	return string(b)
}

// GetClanWars returns a list of wars a clan has particiapted in
func GetClanWars(clanTag string, qparms rest.QParms) ([]ClanWar, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(clanTag))
	sb.WriteString("/warlog")
	url := sb.String()
	log.Trace(url)

	body, err := get(url, qparms)
	if err != nil {
		return nil, err
	}

	// Parse into an array of clans
	type respType struct {
		WarLog []ClanWar `json:"items"`
	}
	var resp respType
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	// Find any "empty" wars and remove it from the list
	for i := range resp.WarLog {
		if i >= len(resp.WarLog) {
			break
		}
		war := resp.WarLog[i]
		if war.Opponent.Name == "" {
			if i == len(resp.WarLog) {
				resp.WarLog = resp.WarLog[:i]
			} else {
				resp.WarLog = append(resp.WarLog[:i], resp.WarLog[i+1:]...)
			}
		}
	}

	// Remove wars without an opponent clan's name
	i := 0
	warLog := make([]ClanWar, len(resp.WarLog))
	for _, war := range resp.WarLog {
		if war.Opponent.Name != "" {
			warLog[i] = war
			i++
		}
	}

	// Return the trimmed war log
	return warLog[:i], nil
}

// GetCurrentWar returns information about the current war a clan is participating in
func GetCurrentWar(clanTag string) (*ClanWar, error) {
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(config.Data.BaseURL)
	sb.WriteString("/clans/")
	sb.WriteString(fmtTag(clanTag))
	sb.WriteString("/currentwar")
	url := sb.String()
	log.Trace(url)

	// Send the request and get the response
	body, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	// Parse into a war
	var war ClanWar
	if err := json.Unmarshal(body, &war); err != nil {
		log.Debug("failed to parse the json response")
		return nil, err
	}

	// Check to see if the clan is in a war
	if war.State == "notInWar" {
		return nil, ErrNotInWar
	}

	return &war, nil
}
