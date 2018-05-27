package worker

import (
	"context"
	"sync"
	"time"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/schema"
	"github.com/alextanhongpin/go-github-scraper/internal/util"
	"github.com/alextanhongpin/go-github-scraper/service/analyticsvc"
	"github.com/alextanhongpin/go-github-scraper/service/profilesvc"
	"github.com/alextanhongpin/go-github-scraper/service/reposvc"
	"github.com/alextanhongpin/go-github-scraper/service/usersvc"

	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	// Worker exposes the interface
	Worker interface {
		NewFetchUsers(tab string) *cron.Cron
		NewFetchRepos(tab string) *cron.Cron
		NewAnalyticBuilder(tab string) *cron.Cron
		NewProfileBuilder(tab string) *cron.Cron
		// NewMatchingBuilder (tab string)*cron.Cron
	}

	worker struct {
		gsvc github.API
		asvc analyticsvc.Service
		psvc profilesvc.Service
		rsvc reposvc.Service
		usvc usersvc.Service
		zlog *zap.Logger
	}
)

// New creates a new worker
func New(gsvc github.API, asvc analyticsvc.Service, psvc profilesvc.Service, rsvc reposvc.Service, usvc usersvc.Service, zlog *zap.Logger) Worker {
	return &worker{
		gsvc: gsvc,
		asvc: asvc,
		usvc: usvc,
		rsvc: rsvc,
		psvc: psvc,
		zlog: zlog,
	}
}

// makeEndDate sets the end date n months away to max present day
func makeEndDate(start string, months int) string {
	t1, err := time.Parse("2006-01-02", start)
	if err != nil {
		return time.Now().Format("2006-01-02")
	}
	t2 := t1.Add(time.Duration(months) * 30 * 24 * time.Hour)
	if t2.Unix() > time.Now().Unix() {
		return time.Now().Format("2006-01-02")
	}
	return t2.Format("2006-01-02")
}

func (w *worker) NewFetchUsers(tab string) *cron.Cron {
	zlog := util.LoggerWithRequestID(w.zlog)

	c := cron.New()
	c.AddFunc(tab, func() {
		zlog.Info("start cron", zap.String("type", "fetch_users"))
		start, ok := w.usvc.FindLastCreated()
		months := 6
		end := makeEndDate(start, months)
		zlog.Info("fetch users since",
			zap.String("start", start),
			zap.String("end", end),
			zap.Bool("default", ok))

		users, err := w.gsvc.FetchUsersCursor(viper.GetString("github_location"), start, end, 30)
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
	zlog := util.LoggerWithRequestID(w.zlog)
	userPerPage := 10
	repoPerPage := 30

	c := cron.New()
	c.AddFunc(tab, func() {
		zlog.Info("start cron", zap.String("type", "fetch_repos"))
		users, err := w.usvc.FindLastFetched(userPerPage)
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

			repos, err := w.gsvc.FetchReposCursor(login, start, end, repoPerPage)
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
	zlog := util.LoggerWithRequestID(w.zlog)
	zlog.Info("NewAnalyticBuilder", zap.Bool("initialized", true))

	c := cron.New()
	c.AddFunc(tab, func() {
		zlog.Info("start cron", zap.String("type", "build_analytic"))

		// Start: user_count
		count, err := w.usvc.Count()
		if err != nil {
			zlog.Warn("error getting user count", zap.Error(err))
		}
		if err := w.asvc.PostUserCount(count); err != nil {
			zlog.Warn("error updating user count", zap.Error(err))
		}
		// End: user_count

		// Start: repo_count
		count, err = w.rsvc.Count()
		if err != nil {
			zlog.Warn("error getting repo count", zap.Error(err))
		}
		if err := w.asvc.PostRepoCount(count); err != nil {
			zlog.Warn("error updating repo count", zap.Error(err))
		}
		// End: repo_count

		// Start: repos_most_recent
		repos, err := w.rsvc.MostRecent(10)
		if err != nil {
			zlog.Warn("error fetching most recent repos", zap.Error(err))
		}

		if err := w.asvc.PostReposMostRecent(repos); err != nil {
			zlog.Warn("error updating most recent repos", zap.Error(err))
		}
		// End: repos_most_recent

		// Start: repo_count_by_user
		users, err := w.rsvc.RepoCountByUser(10)
		if err != nil {
			zlog.Warn("error fetching repo count by users", zap.Error(err))
		}

		if err := w.asvc.PostRepoCountByUser(users); err != nil {
			zlog.Warn("error updating repo count by users", zap.Error(err))
		}
		// End: repo_count_by_user

		// Start: repos_most_stars
		repos, err = w.rsvc.MostStars(10)
		if err != nil {
			zlog.Warn("error fetching most stars repos", zap.Error(err))
		}
		if err := w.asvc.PostReposMostStars(repos); err != nil {
			zlog.Warn("error updating repo count by users", zap.Error(err))
		}
		// End: repos_most_stars

		// Start: languages_most_popular
		languages, err := w.rsvc.MostPopularLanguage(20)
		if err != nil {
			zlog.Warn("error fetching language most popular", zap.Error(err))
		}
		if err := w.asvc.PostMostPopularLanguage(languages); err != nil {
			zlog.Warn("error updating language most popular", zap.Error(err))
		}
		// End: languages_most_popular

		// Start: language_count_by_user
		// languages, err := w.rsvc.LanguageCountByUser("login", 10)
		// if err != nil {
		// 	zlog.Warn("error fetching language count repos", zap.Error(err))
		// }
		// if err := w.asvc.PostLanguageCountByUser(languages); err != nil {
		// 	zlog.Warn("error updating language most popular", zap.Error(err))
		// }
		// End: language_count_by_user

		var reposByLanguages []schema.RepoLanguage
		var userCountByLanguage []schema.UserCountByLanguage
		for _, lang := range languages {
			// Start: repos_most_recent_by_language
			r, err := w.rsvc.MostRecentReposByLanguage(lang.Name, 10)
			if err != nil {
				zlog.Warn("error fetching most recent repos by language", zap.Error(err))
			}
			reposByLanguages = append(reposByLanguages, schema.RepoLanguage{
				Language: lang.Name,
				Repos:    r,
			})
			// End: repos_most_recent_by_language

			// Start: repos_by_language
			users, err := w.rsvc.ReposByLanguage(lang.Name, 10)
			if err != nil {
				zlog.Warn("error fetching user repo count by language", zap.Error(err))
			}
			userCountByLanguage = append(userCountByLanguage, schema.UserCountByLanguage{
				Language: lang.Name,
				Users:    users,
			})
			// End: repos_by_language
		}
		if err := w.asvc.PostMostRecentReposByLanguage(reposByLanguages); err != nil {
			zlog.Warn("error updating most recent repos by language", zap.Error(err))
		}
		if err := w.asvc.PostReposByLanguage(userCountByLanguage); err != nil {
			zlog.Warn("error updating user repo count by language", zap.Error(err))
		}
		zlog.Info("NewAnalyticBuilder", zap.Bool("ran", true))
	})
	return c
}

func (w *worker) NewProfileBuilder(tab string) *cron.Cron {
	zlog := util.LoggerWithRequestID(w.zlog)
	maxWorkers := 16

	c := cron.New()
	c.AddFunc(tab, func() {})

	go func() {

		zlog.Info("start cron", zap.String("type", "build_profile"))

		logins, err := w.rsvc.DistinctLogin()
		if err != nil {
			zlog.Warn("error getting distinct login", zap.Error(err))
			return
		}
		zlog.Info("got distinct logins",
			zap.Int("count", len(logins)))

		buildProfile := func(l string) schema.Profile {
			watchers, err := w.rsvc.WatchersFor(l)
			if err != nil {
				zlog.Warn("error getting watcher count", zap.String("login", l), zap.Error(err))
			}
			stargazers, err := w.rsvc.StargazersFor(l)
			if err != nil {
				zlog.Warn("error getting stargazer count", zap.String("login", l), zap.Error(err))
			}
			forks, err := w.rsvc.ForksFor(l)
			if err != nil {
				zlog.Warn("error getting fork count", zap.String("login", l), zap.Error(err))
			}
			keywords, err := w.rsvc.KeywordsFor(l, 20)
			if err != nil {
				zlog.Warn("error getting keyword count", zap.String("login", l), zap.Error(err))
			}
			zlog.Info("updated profile",
				zap.String("login", l),
				zap.Int64("watchers", watchers),
				zap.Int64("stargazers", stargazers),
				zap.Int64("forks", forks),
				zap.Int("keywords", len(keywords)))

			return schema.Profile{
				Login:      l,
				Watchers:   watchers,
				Stargazers: stargazers,
				Forks:      forks,
				Keywords:   keywords,
			}
		}

		toStream := func(ctx context.Context, args ...string) <-chan string {
			c := make(chan string)
			go func() {
				defer close(c)
				for _, i := range args {
					select {
					case <-ctx.Done():
						return
					case c <- i:
					}
				}
			}()
			return c
		}

		fanIn := func(ctx context.Context, maxWorkers int, in <-chan string) <-chan schema.Profile {
			c := make(chan schema.Profile)

			var wg sync.WaitGroup
			wg.Add(maxWorkers)

			multiplex := func(worker int, in <-chan string) {
				defer wg.Done()
				for i := range in {
					select {
					case <-ctx.Done():
						return
					case c <- buildProfile(i):
					}
				}
			}

			for i := 0; i < maxWorkers; i++ {
				go multiplex(i, in)
			}

			go func() {
				defer close(c)
				wg.Wait()
			}()
			return c
		}

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		var profiles []schema.Profile
		for p := range fanIn(ctx, maxWorkers, toStream(ctx, logins...)) {
			profiles = append(profiles, p)
		}

		if err = w.psvc.BulkUpsert(profiles); err != nil {
			zlog.Warn("error upserting", zap.Error(err))
		}
		zlog.Info("upserted new profile", zap.Int("count", len(profiles)))
	}()
	// Run this in the background

	return c
}
