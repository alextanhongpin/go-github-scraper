package schema

// Keyword represents the map reduce keyword from MongoDB
type Keyword struct {
	ID    string `json:"name,omitempty" bson:"_id"`
	Value int    `json:"count,omitempty" bson:"value"`
}
