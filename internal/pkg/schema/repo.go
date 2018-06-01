package schema

// Repo represents the repository structure
type Repo struct {
	Name          string   `json:"name" bson:"name,omitempty"`
	CreatedAt     string   `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt     string   `json:"updatedAt" bson:"updatedAt,omitempty"`
	FetchedAt     string   `json:"fetchedAt" bson:"fetchedAt,omitempty"`
	Description   string   `json:"description" bson:"description,omitempty"`
	Languages     []string `json:"languages" bson:"languages,omitempty"`
	HomepageURL   string   `json:"homepageUrl" bson:"homepageUrl,omitempty"`
	ForkCount     int64    `json:"forkCount" bson:"forkCount,omitempty"` // TODO: Change this to `forks``
	IsFork        bool     `json:"isFork" bson:"isFork,omitempty"`
	NameWithOwner string   `json:"nameWithOwner" bson:"nameWithOwner,omitempty"`
	Login         string   `json:"login" bson:"login,omitempty"`
	AvatarURL     string   `json:"avatarUrl" bson:"avatarUrl,omitempty"`
	Stargazers    int64    `json:"stargazers" bson:"stargazers,omitempty"`
	Watchers      int64    `json:"watchers" bson:"watchers,omitempty"`
	URL           string   `json:"url" bson:"url,omitempty"`
}
