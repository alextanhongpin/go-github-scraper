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
	Users    []User `json:"users,omitempty" bson:"users,omitempty"`
}

type ReposMostStars struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Repos    []schema.Repo `json:"repos,omitempty" bson:"repos,omitempty"`
}

type MostPopularLanguage struct {
	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo  `bson:",inline"`
	Languages []Language `json:"languages,omitempty" bson:"languages,omitempty"`
}

type LanguageCountByUser struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Users    []User `json:"users,omitempty" bson:"users,omitempty"`
}

type MostRecentReposByLanguage struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Repos    []schema.Repo `json:"repos,omitempty" bson:"repos,omitempty"`
}

type ReposByLanguage struct {
	Type     string `json:"type,omitempty" bson:"type,omitempty"`
	DateInfo `bson:",inline"`
	Users    []User `json:"users,omitempty" bson:"users,omitempty"`
}

// // Repo represents the repository structure, duplicated from the repo service
// type Repo struct {
// 	Name          string   `json:"name,omitempty" bson:"name,omitempty"`
// 	CreatedAt     string   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
// 	UpdatedAt     string   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
// 	FetchedAt     string   `json:"fetchedAt,omitempty" bson:"fetchedAt,omitempty"`
// 	Description   string   `json:"description,omitempty" bson:"description,omitempty"`
// 	Languages     []string `json:"languages,omitempty" bson:"languages,omitempty"`
// 	HomepageURL   string   `json:"homepageUrl,omitempty" bson:"homepageUrl,omitempty"`
// 	ForkCount     int64    `json:"forkCount,omitempty" bson:"forkCount,omitempty"`
// 	IsFork        bool     `json:"isFork,omitempty" bson:"isFork,omitempty"`
// 	NameWithOwner string   `json:"nameWithOwner,omitempty" bson:"nameWithOwner,omitempty"`
// 	Login         string   `json:"login,omitempty" bson:"login,omitempty"`
// 	AvatarURL     string   `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
// 	Stargazers    int64    `json:"stargazers,omitempty" bson:"stargazers,omitempty"`
// 	Watchers      int64    `json:"watchers,omitempty" bson:"watchers,omitempty"`
// 	URL           string   `json:"url,omitempty" bson:"url,omitempty"`
// }

// Language represents the language and the count
type Language struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Count int    `json:"count,omitempty" bson:"count,omitempty"`
}

// User represents the user and the count
type User struct {
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Count     int    `json:"count,omitempty" bson:"count,omitempty"`
	AvatarURL string `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
}
