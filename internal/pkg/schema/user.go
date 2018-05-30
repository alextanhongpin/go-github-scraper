package schema

// User represents the simplified user struct
type User struct {
	Login     string  `json:"login,omitempty" bson:"login,omitempty"`
	Score     float64 `json:"score,omitempty" bson:"score,omitempty"`
	AvatarURL string  `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
}
