package model

import "time"

// User represents the user information in Github
type User struct {
	Name       string    `json:"name,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	Login      string    `json:"login,omitempty"`
	Bio        string    `json:"bio,omitempty"`
	Location   string    `json:"location,omitempty"`
	Email      string    `json:"email,omitempty"`
	Company    string    `json:"company,omitempty"`
	AvatarURL  string    `json:"avatarUrl,omitempty"`
	WebsiteURL string    `json:"websiteUrl,omitempty"`
}
