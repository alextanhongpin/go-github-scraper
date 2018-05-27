package reposvc

// // Repo represents the repository structure
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

// LanguageCount represents the language and the count
type LanguageCount struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Count int    `json:"count,omitempty" bson:"count,omitempty"`
}

// UserCount represents the user and the count
type UserCount struct {
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Count     int    `json:"count,omitempty" bson:"count,omitempty"`
	AvatarURL string `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
}

type GetRepoCountResponse struct {
	Count int `json:"count,omitempty"`
}
