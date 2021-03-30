package coc

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	cocTimeLayout = "20060102T150405.000Z"
)

// CoCTime is a redefinition of the Time structure.  This allows for unmarshalling of
// the time format used by Clash of Clans.
type CoCTime time.Time

// UnmarshalJSON parses a JSON string into a CocTime structure
func (ct *CoCTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(cocTimeLayout, s)
	if err != nil {
		return err
	}
	*ct = CoCTime(t)
	return nil
}

// MarshalJSON converts a CoCTime into a JSON string
func (ct CoCTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct)
}

// Format prints the date
func (ct CoCTime) Format(s string) string {
	return ct.String()
}

// String converts the date to a string
func (ct CoCTime) String() string {
	t := time.Time(ct)
	return t.String()
}
