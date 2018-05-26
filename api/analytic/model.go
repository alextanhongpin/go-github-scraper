package analytic

// Model represents the interface for the analytic business logic
type (
	Model interface {
		GetUserCount() (*UserCount, error)
		PostUserCount(count int) error
	}

	model struct {
		store Store
	}
)

// NewModel returns a new analytic model
func NewModel(s Store) Model {
	return &model{
		store: s,
	}
}

func (m *model) GetUserCount() (*UserCount, error) {
	// Validate...
	// Tracing...
	return m.store.GetUserCount()
}

func (m *model) PostUserCount(count int) error {
	return m.store.PostUserCount(count)
}
