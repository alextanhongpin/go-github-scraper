package profilesvc

import (
	"context"
	"log"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
	"go.uber.org/zap"
)

type model struct {
	store Store
}

// NewModel returns a new model that fulfils the Service interface
func NewModel(store Store) Service {
	m := model{store: store}
	if err := m.Init(context.Background()); err != nil {
		log.Fatal(err)
	}
	return &m
}

// Perform initialization of the service, such as setting up
// tables for the storage or indexes
func (m *model) Init(ctx context.Context) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "Init"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error initializing profilesvc", zap.Error(err))
		} else {
			zlog.Info("initialize profilesvc")
		}
	}(time.Now())
	return m.store.Init()
}

func (m *model) GetProfile(ctx context.Context, login string) (p *schema.Profile, err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "Init"),
				zap.Duration("took", time.Since(start)),
				zap.String("login", login))
		if err != nil {
			zlog.Warn("error getting profile", zap.Error(err))
		} else {
			zlog.Info("got profile")
		}
	}(time.Now())
	return m.store.GetProfile(login)
}

func (m *model) UpdateProfile(ctx context.Context, login string, profile schema.Profile) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "UpdateProfile"),
				zap.Duration("took", time.Since(start)),
				zap.String("login", login))
		if err != nil {
			zlog.Warn("error updating profile", zap.Error(err))
		} else {
			zlog.Info("update profile")
		}
	}(time.Now())
	return m.store.UpdateProfile(login, profile)
}

func (m *model) BulkUpsert(ctx context.Context, profiles []schema.Profile) (err error) {
	defer func(start time.Time) {
		zlog := logger.RequestIDFromContext(ctx).
			With(zap.String("method", "BulkUpsert"),
				zap.Duration("took", time.Since(start)),
				zap.Int("count", len(profiles)))
		if err != nil {
			zlog.Warn("error bulk upserting profile", zap.Error(err))
		} else {
			zlog.Info("bulk upsert profile")
		}
	}(time.Now())
	return m.store.BulkUpsert(profiles)
}
