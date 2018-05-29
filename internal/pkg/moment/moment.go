package moment

import "time"

// NewUTCDate returns a new UTC date in the RFC3339 format
func NewUTCDate() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// NewCurrentFormattedDate returns a new string date in the format YYYY-MM-DD
func NewCurrentFormattedDate() string {
	return time.Now().Format("2006-01-02")
}
