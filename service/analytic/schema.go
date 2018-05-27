package analyticsvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
)

type DateInfo struct {
	UpdatedAt string `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type UserCount struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Count    int `json:"count,omitempty" bson:"count,omitempty"`
}

type RepoCount struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Count    int `json:"count,omitempty" bson:"count,omitempty"`
}

type ReposMostRecent struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Repos    []schema.Repo `json:"repos,omitempty" bson:"repos,omitempty"`
}

type RepoCountByUser struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Users    []schema.UserCount `json:"users,omitempty" bson:"users,omitempty"`
}

type ReposMostStars struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Repos    []schema.Repo `json:"repos,omitempty" bson:"repos,omitempty"`
}

type MostPopularLanguage struct {
	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo  `bson:",inline"`
	Languages []schema.LanguageCount `json:"languages,omitempty" bson:"languages,omitempty"`
}

type LanguageCountByUser struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Users    []schema.UserCount `json:"users,omitempty" bson:"users,omitempty"`
}

type MostRecentReposByLanguage struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Repos    []schema.RepoLanguage `json:"repos,omitempty" bson:"repos,omitempty"`
}

type ReposByLanguage struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Users    []schema.UserCountByLanguage `json:"users,omitempty" bson:"users,omitempty"`
}
