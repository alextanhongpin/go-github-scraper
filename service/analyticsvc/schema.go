package analyticsvc

import (
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
)

// DateInfo represents the created date and updated date, to be embedded
type DateInfo struct {
	UpdatedAt string `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

// UserCount represents the user count analytics result
type UserCount struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Count    int `json:"count,omitempty" bson:"count,omitempty"`
}

// RepoCount represents the repo count analytics result
type RepoCount struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Count    int `json:"count,omitempty" bson:"count,omitempty"`
}

// ReposMostRecent represents the most recent repos analytics result
type ReposMostRecent struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Repos    []schema.Repo `json:"repos,omitempty" bson:"repos,omitempty"`
}

// RepoCountByUser represents the repo count by users analytics result
type RepoCountByUser struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Users    []schema.UserCount `json:"users,omitempty" bson:"users,omitempty"`
}

// ReposMostStars represents the repos with most stars analytics result
type ReposMostStars struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Repos    []schema.Repo `json:"repos,omitempty" bson:"repos,omitempty"`
}

// MostPopularLanguage represents the languages and the corresponding repo count analytics result
type MostPopularLanguage struct {
	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo  `bson:",inline"`
	Languages []schema.LanguageCount `json:"languages,omitempty" bson:"languages,omitempty"`
}

// LanguageCountByUser represents the languages count by user analytics result
type LanguageCountByUser struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Users    []schema.UserCount `json:"users,omitempty" bson:"users,omitempty"`
}

// MostRecentReposByLanguage represents the most recent repos by language analytics result
type MostRecentReposByLanguage struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Repos    []schema.RepoLanguage `json:"repos,omitempty" bson:"repos,omitempty"`
}

// ReposByLanguage represents the most repo counts for users by language analytics result
type ReposByLanguage struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Users    []schema.UserCountByLanguage `json:"users,omitempty" bson:"users,omitempty"`
}
