package schema

type UserCountByLanguage struct {
	Language string      `json:"language,omitempty" bson:"language,omitempty"`
	Users    []UserCount `json:"users,omitempty" bson:"users,omitempty"`
}
