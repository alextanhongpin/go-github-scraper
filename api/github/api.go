package github

// NewAPI returns a new github api
func NewAPI(token, endpoint string) Model {
	return NewModel(NewStore(token, endpoint))
}
