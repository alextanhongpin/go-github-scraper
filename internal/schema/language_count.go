package schema

// LanguageCount represents the language and the count
type LanguageCount struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Count int    `json:"count,omitempty" bson:"count,omitempty"`
}
