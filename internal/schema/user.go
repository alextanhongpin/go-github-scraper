package schema

import (
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/util"
	"gopkg.in/mgo.v2/bson"
)

// User represents the user information in Github
type User struct {
	Name         string       `json:"name,omitempty"`
	CreatedAt    time.Time    `json:"createdAt,omitempty"`
	UpdatedAt    time.Time    `json:"updatedAt,omitempty"`
	Login        string       `json:"login,omitempty"`
	Bio          string       `json:"bio,omitempty"`
	Location     string       `json:"location,omitempty"`
	Email        string       `json:"email,omitempty"`
	Company      string       `json:"company,omitempty"`
	AvatarURL    string       `json:"avatarUrl,omitempty"`
	WebsiteURL   string       `json:"websiteUrl,omitempty"`
	Repositories Repositories `json:"repositories,omitempty"`
	Gists        Gists        `json:"gists,omitempty"`
	Followers    Followers    `json:"followers,omitempty"`
	Following    Following    `json:"following,omitempty"`
}

// BSON returns the bson object of the user
func (u User) BSON() bson.M {
	return bson.M{
		"name":         u.Name,
		"createdAt":    u.CreatedAt,
		"updatedAt":    u.UpdatedAt,
		"fetchedAt":    util.NewUTCDate(), // Additional field
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
