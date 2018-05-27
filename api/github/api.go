package github

// API represents the interface for Github's API
type API interface {
	FetchUsersCursor(location, start, end string, limit int) ([]User, error)
	FetchReposCursor(login, start, end string, limit int) ([]Repo, error)
}

// New returns a new github api
func New(token, endpoint string) API {
	return NewModel(NewStore(token, endpoint))
}
