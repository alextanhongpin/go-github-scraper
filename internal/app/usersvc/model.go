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

func (m *model) Init(ctx context.Context) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "Init"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error initializing usersvc", zap.Error(err))
		} else {
			zlog.Info("initialiaze usersvc")
		}
	}(time.Now())
	return m.store.Init()
}

func (m *model) MostRecent(ctx context.Context, limit int) (users []User, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "MostRecent"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error getting most recent users", zap.Error(err))
		} else {
			zlog.Info("get most recent users")
		}
	}(time.Now())
	return m.store.FindAll(limit, []string{"-updatedAt"})
}

func (m *model) BulkUpsert(ctx context.Context, users []github.User) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "BulkUpsert"),
				zap.Duration("took", time.Since(start)),
				zap.Int("count", len(users)))
		if err != nil {
			zlog.Warn("error upserting bulk users", zap.Error(err))
		} else {
			zlog.Info("upsert bulk users")
		}
	}(time.Now())
	return m.store.BulkUpsert(users)
}

func (m *model) Drop(ctx context.Context) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "Drop"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error dropping user collection", zap.Error(err))
		} else {
			zlog.Info("drop user collection")
		}
	}(time.Now())
	return m.store.Drop()
}

// FindLastCreated returns the last created date in the format YYYY-MM-DD, and a boolean to indicate
// if the value returned exists or is default
func (m *model) FindLastCreated(ctx context.Context) (lastCreated string, ok bool) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "FindLastCreated"),
				zap.Duration("took", time.Since(start)))
		zlog.Info("find last created",
			zap.String("date", lastCreated),
			zap.Bool("default", ok))
	}(time.Now())
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

func (m *model) FindLastFetched(ctx context.Context, limit int) (res []User, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "FindLastFetched"),
				zap.Duration("took", time.Since(start)),
				zap.Int("limit", limit))
		if err != nil {
			zlog.Warn("error finding last fetched", zap.Error(err))
		} else {
			zlog.Info("find last fetched")
		}
	}(time.Now())
	return m.store.FindAll(limit, []string{"-fetchedAt"})
}

func (m *model) Count(ctx context.Context) (count int, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "Count"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error getting user count", zap.Error(err))
		} else {
			zlog.Info("get user count", zap.Int("count", count))
		}
	}(time.Now())
	return m.store.Count()
}

func (m *model) UpdateOne(ctx context.Context, login string) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "UpdateOne"),
				zap.Duration("took", time.Since(start)),
				zap.String("login", login))
		if err != nil {
			zlog.Warn("error updating one user", zap.Error(err))
		} else {
			zlog.Info("update one user")
		}
	}(time.Now())
	return m.store.UpdateOne(login)
}

func (m *model) FindOne(ctx context.Context, login string) (res *User, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "FindOne"),
				zap.Duration("took", time.Since(start)),
				zap.String("login", login))
		if err != nil {
			zlog.Warn("error finding one user", zap.Error(err))
		} else {
			zlog.Info("find one user")
		}
	}(time.Now())
	return m.store.FindOne(login)
}

func (m *model) PickLogin(ctx context.Context) (logins []string, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "PickLogin"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error picking logins", zap.Error(err))
		} else {
			zlog.Info("pick logins", zap.Int("count", len(logins)))
		}
	}(time.Now())
	return m.store.PickLogin()
}