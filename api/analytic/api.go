package analytic

// API represents the business logic for the analytic service
type API interface {
	GetUserCount() (*UserCount, error)
	PostUserCount(count int) error
}

type api struct {
	store Store
}

// NewAPI returns a new analytic model
func NewAPI(s Store) API {
	return &api{
		store: s,
	}
}

func (a *api) GetUserCount() (*UserCount, error) {
	return a.store.GetUserCount()
}

func (a *api) PostUserCount(count int) error {
	return a.store.PostUserCount(count)
}
