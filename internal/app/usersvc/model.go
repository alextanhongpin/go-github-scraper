package usersvc

import (
	"context"
	"log"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/constant"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"go.uber.org/zap"
)

type model struct {
	store Store
}

// NewModel returns a new model with the store
func NewModel(store Store) Service {
	m := model{store: store}
	if err := m.Init(context.Background()); err != nil {
		log.Fatal(err)
	}
	return &m
}

func (m *model) Init(ctx context.Context) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Info("init")
	return m.store.Init()
}

func (m *model) MostRecent(ctx context.Context, limit int) ([]User, error) {
	return m.store.FindAll(limit, []string{"-updatedAt"})
}

func (m *model) BulkUpsert(ctx context.Context, users []github.User) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Info("BulkUpsert",
		zap.Bool("start", true),
		zap.Int("count", len(users)))
	return m.store.BulkUpsert(users)
}

func (m *model) Drop(ctx context.Context) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Warn("Drop")
	return m.store.Drop()
}

// FindLastCreated returns the last created date in the format YYYY-MM-DD, and a boolean to indicate
// if the value returned exists or is default
func (m *model) FindLastCreated(ctx context.Context) (string, bool) {
	user, err := m.store.FindLastCreated()
	if err != nil || user == nil {
		return constant.GithubCreatedAt, false
	}
	t, err := time.Parse(time.RFC3339, user.CreatedAt)
	if err != nil {
		return constant.GithubCreatedAt, false
	}
	return t.Format("2006-01-02"), true
}

func (m *model) FindLastFetched(ctx context.Context, limit int) ([]User, error) {
	return m.store.FindAll(limit, []string{"-fetchedAt"})
}

func (m *model) Count(ctx context.Context) (int, error) {
	return m.store.Count()
}

func (m *model) UpdateOne(ctx context.Context, login string) error {
	zlog := logger.RequestIDFromContext(ctx)
	zlog.Info("UpdateOne",
		zap.Bool("start", true),
		zap.String("for", login))
	return m.store.UpdateOne(login)
}

func (m *model) FindOne(ctx context.Context, login string) (*User, error) {
	return m.store.FindOne(login)
}

func (m *model) PickLogin(ctx context.Context) ([]string, error) {
	return m.store.PickLogin()
}
