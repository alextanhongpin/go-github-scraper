package cronjob

import "github.com/robfig/cron"

type (
	// CronJob exposes the interface
	CronJob interface {
		Do()
	}

	// Config represents the cronjob config
	Config struct {
		Name        string
		Description string
		Start       bool
		CronTab     string
		Fn          func() error
	}
)

// Do will create a new cron job with the provided configuration
func (cfg *Config) Do() {
	if !cfg.Start {
		return
	}
	c := cron.New()
	c.AddFunc(cfg.CronTab, func() {
		cfg.Fn()
	})
	c.Start()
}

// Exec runs a list of cronjobs
func Exec(jobs ...CronJob) {
	for _, j := range jobs {
		j.Do()
	}
}
