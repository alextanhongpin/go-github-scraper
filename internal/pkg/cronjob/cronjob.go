package cronjob

import (
	"context"

	"github.com/robfig/cron"
	"go.uber.org/zap"
)

type (
	// CronJob exposes the interface
	CronJob interface {
		Do(ctx context.Context)
	}

	// Config represents the cronjob config
	Config struct {
		Name        string
		Description string
		Start       bool
		CronTab     string
		Trigger     bool
		Fn          func(ctx context.Context) error
	}
)

// Do will create a new cron job with the provided configuration
func (cfg *Config) Do(ctx context.Context) {

	if cfg.Trigger {
		go cfg.Fn(ctx)
	}

	if !cfg.Start {
		return
	}
	c := cron.New()
	c.AddFunc(cfg.CronTab, func() {
		cfg.Fn(ctx)
	})
	c.Start()
	zap.L().Info("started cron",
		zap.String("name", cfg.Name),
		zap.String("tab", cfg.CronTab))
}

// Exec runs a list of cronjobs
func Exec(ctx context.Context, jobs ...CronJob) {
	for _, j := range jobs {
		j.Do(ctx)
	}
}
