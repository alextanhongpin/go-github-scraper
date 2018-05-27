package worker

import (
	"time"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/util"
	"github.com/alextanhongpin/go-github-scraper/service/analytic"
	"github.com/alextanhongpin/go-github-scraper/service/repo"
	"github.com/alextanhongpin/go-github-scraper/service/user"

	"github.com/robfig/cron"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type (
	// Worker exposes the interface
	Worker interface {
		NewFetchUsers(tab string) *cron.Cron
		NewFetchRepos(tab string) *cron.Cron
		NewAnalyticBuilder(tab string) *cron.Cron
	}

	worker struct {
		gsvc github.API
		asvc analyticsvc.Service
		usvc usersvc.Service
		rsvc reposvc.Service
		zlog *zap.Logger
	}
)

// New creates a new worker
func New(gsvc github.API, asvc analyticsvc.Service, usvc usersvc.Service, rsvc reposvc.Service, zlog *zap.Logger) Worker {
	return &worker{
		gsvc: gsvc,
		asvc: asvc,
		usvc: usvc,
		rsvc: rsvc,
		zlog: zlog,
		// config: config
	}
}

func (w *worker) NewFetchUsers(tab string) *cron.Cron {
	requestID, err := uuid.NewV4()
	zlog := w.zlog
	if err != nil {
		zlog.Warn("error generating uuid", zap.Error(err))
	} else {
		zlog = w.zlog.WithOptions(zap.Fields(zap.String("requestId", requestID.String())))
	}

	c := cron.New()
	c.AddFunc(tab, func() {
		start, ok := w.usvc.FindLastCreated()
		t, err := time.Parse("2006-01-02", start)
		if err != nil {
			zlog.Warn("error parsing time", zap.Error(err))
		}
		t2 := t.Add(6 * 30 * 24 * time.Hour)
		var end string
		if t2.Unix() > time.Now().Unix() {
			end = time.Now().Format("2006-01-02")
		} else {
			end = t2.Format("2006-01-02")
		}
		zlog.Info("fetch users since",
			zap.String("start", start),
			zap.String("end", end),
			zap.Bool("default", ok))

		// TODO: Make Malaysia a config
		users, err := w.gsvc.FetchUsersCursor("Malaysia", start, end, 30)
		if err != nil {
			zlog.Warn("error fetching users", zap.Error(err))
		}

		if err := w.usvc.BulkUpsert(users); err != nil {
			zlog.Warn("error upserting users", zap.Error(err))
		}

		zlog.Info("upserted users",
			zap.String("start", start),
			zap.String("end", end),
			zap.Int("count", len(users)))
	})
	return c
}

func (w *worker) NewFetchRepos(tab string) *cron.Cron {
	requestID, err := uuid.NewV4()
	zlog := w.zlog
	if err != nil {
		zlog.Warn("error generating uuid", zap.Error(err))
	} else {
		zlog = w.zlog.WithOptions(zap.Fields(zap.String("requestId", requestID.String())))
	}

	c := cron.New()
	c.AddFunc(tab, func() {
		// TODO: Config
		users, err := w.usvc.FindLastFetched(10)
		if err != nil {
			zlog.Error("unable to find last fetched users",
				zap.Error(err))
			return
		}

		for _, user := range users {
			login := user.Login
			if login == "" {
				continue
			}
			start, ok := w.rsvc.FindLastCreatedByUser(login)
			end := util.NewCurrentFormattedDate()
			zlog.Info("fetch repos since",
				zap.String("start", start),
				zap.String("end", end),
				zap.String("for", login),
				zap.Bool("default", ok))

			// TODO: Make Malaysia a config
			repos, err := w.gsvc.FetchReposCursor(login, start, end, 30)
			if err != nil {
				zlog.Warn("error fetching repos", zap.Error(err))
			}

			if err := w.rsvc.BulkUpsert(repos); err != nil {
				zlog.Warn("error upserting repos", zap.Error(err))
			}

			if err = w.usvc.UpdateOne(login); err != nil {
				zlog.Warn("error updating users", zap.Error(err))
			}

			zlog.Info("upserted repos",
				zap.Int("count", len(repos)),
				zap.String("start", start),
				zap.String("end", end),
				zap.String("for", login))
		}
	})
	return c
}

func (w *worker) NewAnalyticBuilder(tab string) *cron.Cron {
	requestID, err := uuid.NewV4()
	zlog := w.zlog
	if err != nil {
		zlog.Warn("error generating uuid", zap.Error(err))
	} else {
		zlog = w.zlog.WithOptions(zap.Fields(zap.String("requestId", requestID.String())))
	}

	zlog.Info("NewAnalyticBuilder", zap.Bool("initialized", true))

	c := cron.New()
	c.AddFunc(tab, func() {
		count, err := w.usvc.Count()
		if err != nil {
			zlog.Warn("error getting user count", zap.Error(err))
		}
		if err := w.asvc.PostUserCount(count); err != nil {
			zlog.Warn("error updating user count", zap.Error(err))
		}

		count, err = w.rsvc.Count()
		if err != nil {
			zlog.Warn("error getting repo count", zap.Error(err))
		}
		if err := w.asvc.PostRepoCount(count); err != nil {
			zlog.Warn("error updating repo count", zap.Error(err))
		}

		repos, err := w.rsvc.MostRecent(10)
		if err != nil {
			zlog.Warn("error fetching most recent repos", zap.Error(err))
		}

		if err := w.asvc.PostReposMostRecent(repos); err != nil {
			zlog.Warn("error updating most recent repos", zap.Error(err))
		}

		// repos, err = w.rsvc.MostStars(10)
		// if err != nil {
		// 	zlog.Warn("error fetching most stars repos", zap.Error(err))
		// }

		// languages, err := w.rsvc.MostPopularLanguage(20)
		// if err != nil {
		// 	zlog.Warn("error fetching language count repos", zap.Error(err))
		// }

		// users, err := w.rsvc.RepoCountByUser(10)
		// if err != nil {
		// 	zlog.Warn("error fetching language count repos", zap.Error(err))
		// }

		// languages, err := w.rsvc.LanguageCountByUser("login", 10)
		// if err != nil {
		// 	zlog.Warn("error fetching language count repos", zap.Error(err))
		// }

		// for _, lang := range languages {
		// 	repos, err := w.rsvc.MostRecentReposByLanguage(lang.Name, 20)
		// 	if err != nil {
		// 		zlog.Warn("error fetching language count repos", zap.Error(err))
		// 	}

		// 	users, err := w.rsvc.ReposByLanguage(lang.Name, 10)
		// 	if err != nil {
		// 		zlog.Warn("error fetching language count repos", zap.Error(err))
		// 	}
		// }

		// Count() (int, error)
		// MostRecent(limit int) ([]Repo, error)
		// MostStars(limit int) ([]Repo, error)
		// MostPopularLanguage(limit int) ([]LanguageCount, error)
		// RepoCountByUser(limit int) ([]UserCount, error)
		// LanguageCountByUser(login string, limit int) ([]LanguageCount, error)
		// MostRecentReposByLanguage(language string, limit int) ([]Repo, error)
		// ReposByLanguage(language string, limit int) ([]UserCount, error)
		zlog.Info("NewAnalyticBuilder", zap.Bool("ran", true))
	})
	return c
}
