package schema

type RepoLanguage struct {
	Language string `json:"language,omitempty" bson:"language,omitempty"`
	Repos    []Repo `json:"repos,omitempty" bson:"repos,omitempty"`
}
