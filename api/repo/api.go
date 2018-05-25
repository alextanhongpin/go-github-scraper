package repo

type API interface {
	MostRecent() ([]Repo, error)
}

type api struct {
	store Store
}

func NewAPI(store Store) API {
	return &api{
		store: store,
	}
}

func (a *api) MostRecent() ([]Repo, error) {
	return a.store.FindAll(20, []string{"-updatedAt"})
}

func (a *api) MostStars() ([]Repo, error) {
	return a.store.FindAll(20, []string{"-stargazers"})
}
