package analytic

import "time"

// UserCount represents the `user_count` analytic type
type UserCount struct {
	Type      string    `json:"type,omitempty"`
	Count     int       `json:"count,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
