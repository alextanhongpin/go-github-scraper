package github

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Repo represents the repository structure
type Repo struct {
	Name          string     `json:"name,omitempty" bson:"name,omitempty"`
	CreatedAt     time.Time  `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time  `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Description   string     `json:"description,omitempty" bson:"description,omitempty"`
	Languages     Language   `json:"languages,omitempty" bson:"languages,omitempty"`
	HomepageURL   string     `json:"homepageUrl,omitempty" bson:"homepageUrl,omitempty"`
	ForkCount     int64      `json:"forkCount,omitempty" bson:"forkCount,omitempty"`
	IsFork        bool       `json:"isFork,omitempty" bson:"isFork,omitempty"`
	NameWithOwner string     `json:"nameWithOwner,omitempty" bson:"nameWithOwner,omitempty"`
	Owner         Owner      `json:"owner,omitempty" bson:"owner,omitempty"`
	Stargazers    Stargazers `json:"stargazers,omitempty" bson:"stargazers,omitempty"`
	Watchers      Watchers   `json:"watchers,omitempty" bson:"watchers,omitempty"`
	URL           string     `json:"url,omitempty" bson:"url,omitempty"`
}

// BSON returns the repo as bson object
func (r Repo) BSON() bson.M {
	var languages []string
	for _, lang := range r.Languages.Edges {
		languages = append(languages, lang.Node.Name)
	}
	return bson.M{
		"name":          r.Name,
		"createdAt":     r.CreatedAt.UTC().Format(time.RFC3339),
		"updatedAt":     r.UpdatedAt.UTC().Format(time.RFC3339),
		"description":   r.Description,
		"languages":     languages,
		"homepageUrl":   r.HomepageURL,
		"isFork":        r.IsFork,
		"forks":         r.ForkCount,
		"nameWithOwner": r.NameWithOwner,
		"stargazers":    r.Stargazers.TotalCount,
		"watchers":      r.Watchers.TotalCount,
		"login":         r.Owner.Login,
		"avatarUrl":     r.Owner.AvatarURL,
		"url":           r.URL,
	}
}

// Language represents the language node of the repo
type Language struct {
	TotalCount int64          `json:"totalCount,omitempty"`
	Edges      []LanguageEdge `json:"edges,omitempty"`
}

// LanguageEdge represents the language edge of the repo
type LanguageEdge struct {
	Node LanguageNode `json:"node,omitempty"`
}

// LanguageNode represents the language node of the edge
type LanguageNode struct {
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

// Owner represents the repo's owner
type Owner struct {
	Login     string `json:"login,omitempty"`
	AvatarURL string `json:"avatarUrl,omitempty"`
}

// Stargazers represents the stargazers
type Stargazers struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Watchers represents the watchers
type Watchers struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}
