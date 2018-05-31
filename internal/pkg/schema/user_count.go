package schema

// UserCount represents the user and the count
type UserCount struct {
	Name      string `json:"login,omitempty" bson:"name,omitempty"`
	Count     int    `json:"count,omitempty" bson:"count,omitempty"`
	AvatarURL string `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
}
