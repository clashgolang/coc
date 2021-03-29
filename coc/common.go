package coc

import (
	"encoding/json"
	"net/url"
)

// BadgeUrls are the URLs for badges
type BadgeUrls struct {
	Small  string `json:"small"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

// String returns a string representation of a label
func (urls BadgeUrls) String() string {
	b, _ := json.Marshal(urls)
	return string(b)
}

// IconUrls are the URLs for icons
type IconUrls struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
}

// String returns a string representation of a label
func (urls IconUrls) String() string {
	b, _ := json.Marshal(urls)
	return string(b)
}

// fmtTag formats the tag for use in a URL or in a query parameter.
func fmtTag(tag string) string {
	// If the tag doesn't have a '#' character at the front, add one
	if len(tag) == 0 {
		return tag
	} else if tag[0] == '#' {
		return url.QueryEscape(tag)
	} else {
		return url.QueryEscape("#" + tag)
	}
}
