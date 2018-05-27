package reposvc

type GetRepoCountResponse struct {
	Count int `json:"count,omitempty"`
}

type Watchers struct {
	Count int64 `json:"count,omitempty" bson:"count,omitempty"`
}

type Stargazers struct {
	Count int64 `json:"count,omitempty" bson:"count,omitempty"`
}

type Forks struct {
	Count int64 `json:"count,omitempty" bson:"count,omitempty"`
}
