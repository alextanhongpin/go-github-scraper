package schema

import (
	"github.com/alextanhongpin/go-github-scraper/internal/util"
	"gopkg.in/mgo.v2/bson"
)

// Profile represents the Github user with additional metadata
type Profile struct {
	Login      string    `json:"login,omitempty" bson:"login,omitempty"`
	Watchers   int64     `json:"watchers,omitempty" bson:"watchers,omitempty"`
	Stargazers int64     `json:"stargazers,omitempty" bson:"stargazers,omitempty"`
	Forks      int64     `json:"forks,omitempty" bson:"forks,omitempty"`
	Keywords   []Keyword `json:"keywords,omitempty" bson:"keywords,omitempty"`
	Matches    []User    `json:"matches,omitempty" bson:"matches,omitempty"`
	UpdatedAt  string    `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	CreatedAt  string    `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

// BSON converts the profile to MongoDB BSON object
func (p *Profile) BSON() bson.M {
	return bson.M{
		"watchers":   p.Watchers,
		"stargazers": p.Stargazers,
		"forks":      p.Forks,
		"keywords":   p.Keywords,
		"matches":    p.Matches,
		"updatedAt":  util.NewUTCDate(),
	}
}
