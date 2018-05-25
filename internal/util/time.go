package util

import "time"

// NewUTCDate returns a new UTC date in the RFC3339 format
func NewUTCDate() string {
	return time.Now().UTC().Format(time.RFC3339)
}
