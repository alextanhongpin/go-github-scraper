package user

// User represents the user information in Github
type User struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	FetchedAt    string `json:"fetchedAt,omitempty" bson:"fetchedAt,omitempty"`
	Login        string `json:"login,omitempty" bson:"login,omitempty"`
	Bio          string `json:"bio,omitempty" bson:"bio,omitempty"`
	Location     string `json:"location,omitempty" bson:"location,omitempty"`
	Email        string `json:"email,omitempty" bson:"email,omitempty"`
	Company      string `json:"company,omitempty" bson:"company,omitempty"`
	AvatarURL    string `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
	WebsiteURL   string `json:"websiteUrl,omitempty" bson:"websiteUrl,omitempty"`
	Repositories int64  `json:"repositories,omitempty" bson:"repositories,omitempty"`
	Gists        int64  `json:"gists,omitempty" bson:"gists,omitempty"`
	Followers    int64  `json:"followers,omitempty" bson:"followers,omitempty"`
	Following    int64  `json:"following,omitempty" bson:"following,omitempty"`
}
