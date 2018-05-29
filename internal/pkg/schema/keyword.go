package schema

// Keyword represents the map reduce keyword from MongoDB
type Keyword struct {
	ID    string `json:"_id,omitempty" bson:"_id"`
	Value int    `json:"value,omitempty"`
}
