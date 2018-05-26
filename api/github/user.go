package github

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User represents the user information in Github
type User struct {
	Name         string       `json:"name,omitempty" bson:"name,omitempty"`
	CreatedAt    time.Time    `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    time.Time    `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Login        string       `json:"login,omitempty" bson:"login,omitempty"`
	Bio          string       `json:"bio,omitempty" bson:"bio,omitempty"`
	Location     string       `json:"location,omitempty" bson:"location,omitempty"`
	Email        string       `json:"email,omitempty" bson:"email,omitempty"`
	Company      string       `json:"company,omitempty" bson:"company,omitempty"`
	AvatarURL    string       `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
	WebsiteURL   string       `json:"websiteUrl,omitempty" bson:"websiteUrl,omitempty"`
	Repositories Repositories `json:"repositories,omitempty" bson:"repositories,omitempty"`
	Gists        Gists        `json:"gists,omitempty" bson:"gists,omitempty"`
	Followers    Followers    `json:"followers,omitempty" bson:"followers,omitempty"`
	Following    Following    `json:"following,omitempty" bson:"following,omitempty"`
}

// BSON returns the bson object of the user
func (u User) BSON() bson.M {
	return bson.M{
		"name":         u.Name,
		"createdAt":    u.CreatedAt.UTC().Format(time.RFC3339),
		"updatedAt":    u.UpdatedAt.UTC().Format(time.RFC3339),
		"login":        u.Login,
		"bio":          u.Bio,
		"location":     u.Location,
		"email":        u.Email,
		"company":      u.Company,
		"avatarUrl":    u.AvatarURL,
		"websiteUrl":   u.WebsiteURL,
		"repositories": u.Repositories.TotalCount,
		"gists":        u.Gists.TotalCount,
		"followers":    u.Followers.TotalCount,
		"following":    u.Following.TotalCount,
	}
}

// Repositories represent the repository count
type Repositories struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Gists represent the gist count
type Gists struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Followers represent the follower count
type Followers struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Following represent the follower count
type Following struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}
