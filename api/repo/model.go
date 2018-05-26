package repo

// Model represents the interface for the repo business logic
type Model interface {
	MostRecent() ([]Repo, error)
	MostStars() ([]Repo, error)
	// MostPopularLanguage() ([]Language, error)
	// MostRepos() ([]User, error)
}

type model struct {
	store Store
}

// NewModel returns a pointer to the Model
func NewModel(store Store) Model {
	return &model{
		store: store,
	}
}

func (m *model) MostRecent() ([]Repo, error) {
	return m.store.FindAll(20, []string{"-updatedAt"})
}

func (m *model) MostStars() ([]Repo, error) {
	return m.store.FindAll(20, []string{"-stargazers"})
}

// func (m *model) MostPopularLanguage() ([]Language{}, error) {
// 	return m.store.AggregateLanguages(20)
// }

// func (m *model) MostRepos() ([]User, error) {
// 	return m.store.AggregateUsers(20)
// }
