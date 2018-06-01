package schema

import (
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/moment"
	"gopkg.in/mgo.v2/bson"
)

// Profile represents the Github user with additional metadata
type Profile struct {
	Watchers   int64           `json:"watchers,omitempty" bson:"watchers,omitempty"`
	Stargazers int64           `json:"stargazers,omitempty" bson:"stargazers,omitempty"`
	Forks      int64           `json:"forks,omitempty" bson:"forks,omitempty"`
	Languages  []LanguageCount `json:"languages,omitempty" bson:"languages,omitempty"`
	Keywords   []Keyword       `json:"keywords,omitempty" bson:"keywords,omitempty"`
	Matches    []User          `json:"matches,omitempty" bson:"matches,omitempty"`
}

// BSON converts the profile to MongoDB BSON object
func (p *Profile) BSON() bson.M {
	m := bson.M{
		"watchers":   p.Watchers,
		"stargazers": p.Stargazers,
		"forks":      p.Forks,
		"keywords":   p.Keywords,
		"languages":  p.Languages,
		"matches":    p.Matches,
		"updatedAt":  moment.NewUTCDate(),
	}

	if len(p.Keywords) == 0 {
		delete(m, "keywords")
	}

	if len(p.Languages) == 0 {
		delete(m, "languages")
	}

	if len(p.Matches) == 0 {
		delete(m, "matches")
	}
	return m
}
