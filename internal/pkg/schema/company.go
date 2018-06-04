package schema

type Company struct {
	Count   int    `json:"count,omitempty" bson:"count,omitempty"`
	Company string `json:"company,omitempty" bson:"company,omitempty"`
	Users   []User `json:"users,omitempty" bson:"users,omitempty"`
}
