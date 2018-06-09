package reposvc

// GetRepoCountResponse represents the response of the count
// type GetRepoCountResponse struct {
// 	Count int `json:"count"` // Should default to zero if not found
// }

type Watchers struct {
	Count int64 `json:"count,omitempty" bson:"count,omitempty"`
}

type Stargazers struct {
	Count int64 `json:"count,omitempty" bson:"count,omitempty"`
}

type Forks struct {
	Count int64 `json:"count,omitempty" bson:"count,omitempty"`
}
