// Package mediatorsvc handles multiple service orchestrations
// and is categorized as domain model
package mediatorsvc

import (
	"context"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/app/reposvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/statsvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/moment"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"

	"go.uber.org/zap"
)

type (
	// Service represents the methods the mediator service must implement
	Service interface {
		FetchUsers(ctx context.Context, location string, months int, perPage int) error
		FetchRepos(ctx context.Context, userPerPage, repoPerPage int) error
		UpdateUserCount(ctx context.Context) error
		UpdateRepoCount(ctx context.Context) error
		UpdateReposMostRecent(ctx context.Context, perPage int) error
		UpdateRepoCountByUser(ctx context.Context, perPage int) error
		UpdateReposMostStars(ctx context.Context, perPage int) error
		UpdateLanguagesMostPopular(ctx context.Context, perPage int) error
		UpdateMostRecentReposByLanguage(ctx context.Context, perPage int) error
		UpdateReposByLanguage(ctx context.Context, perPage int) error
		UpdateProfile(ctx context.Context, numWorkers int) error
		UpdateMatches(ctx context.Context) error
	}

	// Mediator holds the services in used
	Mediator struct {
		Stat   statsvc.Service
		Github github.API
		// Profile profilesvc.Service
		Repo reposvc.Service
		User usersvc.Service
	}

	model struct {
		Mediator
		logger *logger.Logger
	}
)

// New returns a new mediator service
func New(m Mediator, l *logger.Logger) Service {
	return &model{m, l}
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

func (m *model) FetchUsers(ctx context.Context, location string, months int, perPage int) (err error) {
	var users []github.User
	var start, end string

	defer func(s time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "FetchUsers"),
				zap.Duration("took", time.Since(s)),
				zap.String("location", location),
				zap.Int("months", months),
				zap.Int("perPage", perPage))
		if err != nil {
			zlog.Warn("error fetching and bulk upserting users", zap.Error(err))
		} else {
			zlog.Info("upsert users", zap.Int("count", len(users)))
		}
	}(time.Now())

	start, _ = m.User.FindLastCreated(ctx)
	end = makeEndDate(start, months)

	users, err = m.Github.FetchUsersCursor(ctx, location, start, end, perPage)
	if err != nil {
		return
	}

	if err = m.User.BulkUpsert(ctx, users); err != nil {
		return
	}
	return
}

func (m *model) FetchRepos(ctx context.Context, userPerPage, repoPerPage int) (err error) {
	var users []usersvc.User
	var repos []github.Repo
	defer func(s time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "FetchRepos"),
				zap.Duration("took", time.Since(s)),
				zap.Int("userPerPage", userPerPage),
				zap.Int("repoPerPage", repoPerPage))
		if err != nil {
			zlog.Warn("error fetching and upserting repos", zap.Error(err))
		} else {
			zlog.Info("got users and upsert users",
				zap.Int("users", len(users)),
				zap.Int("repos", len(repos)))
		}
	}(time.Now())

	users, err = m.User.FindLastFetched(ctx, userPerPage)
	if err != nil {
		return
	}

	for _, user := range users {
		login := user.Login
		if login == "" {
			continue
		}

		start, _ := m.Repo.FindLastCreatedByUser(ctx, login)
		end := moment.NewCurrentFormattedDate()

		repos, err = m.Github.FetchReposCursor(ctx, login, start, end, repoPerPage)
		if err != nil {
			return
		}

		if err = m.Repo.BulkUpsert(ctx, repos); err != nil {
			return
		}

		if err = m.User.UpdateOne(ctx, login); err != nil {
			return
		}
	}
	return
}

// UpdateUserCount updates the analytic type `user_count`
func (m *model) UpdateUserCount(ctx context.Context) (err error) {
	var count int
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateUserCount"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error updating user count", zap.Error(err))
		} else {
			zlog.Info("update user count", zap.Int("count", count))
		}
	}(time.Now())

	count, err = m.User.Count(ctx)
	if err != nil {
		return
	}

	if err = m.Stat.PostUserCount(ctx, count); err != nil {
		return
	}
	return
}

// UpdateRepoCount updates the analytic type `repo_count`
func (m *model) UpdateRepoCount(ctx context.Context) (err error) {
	var count int
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateRepoCount"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error updating repo count", zap.Error(err))
		} else {
			zlog.Info("update repo count", zap.Int("count", count))
		}
	}(time.Now())

	count, err = m.Repo.Count(ctx)
	if err != nil {
		return err
	}
	if err = m.Stat.PostRepoCount(ctx, count); err != nil {
		return err
	}
	return
}

// UpdateReposMostRecent updates the analytic type `repos_most_recent`
func (m *model) UpdateReposMostRecent(ctx context.Context, perPage int) (err error) {
	var repos []schema.Repo
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateReposMostRecent"),
				zap.Duration("took", time.Since(start)))
		if err != nil {
			zlog.Warn("error updating repos most recent", zap.Error(err))
		} else {
			zlog.Info("update repo most recent", zap.Int("count", len(repos)))
		}
	}(time.Now())

	repos, err = m.Repo.MostRecent(ctx, perPage)
	if err != nil {
		return
	}

	if err = m.Stat.PostReposMostRecent(ctx, repos); err != nil {
		return
	}
	return
}

// UpdateRepoCountByUser updates the analytic type `repo_count_by_user`
func (m *model) UpdateRepoCountByUser(ctx context.Context, perPage int) (err error) {
	var users []schema.UserCount
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateRepoCountByUser"),
				zap.Duration("took", time.Since(start)),
				zap.Int("perPage", perPage))
		if err != nil {
			zlog.Warn("error updating repo count by user", zap.Error(err))
		} else {
			zlog.Info("update user repo count by user", zap.Int("count", len(users)))
		}
	}(time.Now())

	users, err = m.Repo.RepoCountByUser(ctx, perPage)
	if err != nil {
		return
	}

	if err = m.Stat.PostRepoCountByUser(ctx, users); err != nil {
		return
	}
	return
}

// UpdateReposMostStars updates the analytic type `repos_most_stars`
func (m *model) UpdateReposMostStars(ctx context.Context, perPage int) (err error) {
	var repos []schema.Repo
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateReposMostStars"),
				zap.Duration("took", time.Since(start)),
				zap.Int("perPage", perPage))
		if err != nil {
			zlog.Warn("error updating repos most stars", zap.Error(err))
		} else {
			zlog.Info("update repos most stars", zap.Int("count", len(repos)))
		}
	}(time.Now())

	repos, err = m.Repo.MostStars(ctx, perPage)
	if err != nil {
		return
	}

	if err = m.Stat.PostReposMostStars(ctx, repos); err != nil {
		return
	}
	return
}

// UpdateLanguagesMostPopular updates the analytic type `languages_most_popular`
func (m *model) UpdateLanguagesMostPopular(ctx context.Context, perPage int) (err error) {
	var languages []schema.LanguageCount
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateLanguagesMostPopular"),
				zap.Duration("took", time.Since(start)),
				zap.Int("perPage", perPage))
		if err != nil {
			zlog.Warn("error updating language most popular", zap.Error(err))
		} else {
			zlog.Info("update language most popular", zap.Int("count", len(languages)))
		}
	}(time.Now())

	languages, err = m.Repo.MostPopularLanguage(ctx, perPage)
	if err != nil {
		return
	}

	if err = m.Stat.PostMostPopularLanguage(ctx, languages); err != nil {
		return
	}
	return
}

// UpdateMostRecentReposByLanguage updates the analytic type `repos_most_recent_by_language`
func (m *model) UpdateMostRecentReposByLanguage(ctx context.Context, perPage int) (err error) {
	var languages []schema.LanguageCount
	var r []schema.Repo
	var repos []schema.RepoLanguage

	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateMostRecentReposByLanguage"),
				zap.Duration("took", time.Since(start)),
				zap.Int("perPage", perPage))
		if err != nil {
			zlog.Warn("error updating most recent repos by language", zap.Error(err))
		} else {
			zlog.Info("update most recent repos by language", zap.Int("count", len(repos)))
		}
	}(time.Now())

	languages, err = m.Repo.MostPopularLanguage(ctx, perPage)
	if err != nil {
		return err
	}

	for _, lang := range languages {
		r, err = m.Repo.MostRecentReposByLanguage(ctx, lang.Name, perPage)
		if err != nil {
			return
		}
		repos = append(repos, schema.RepoLanguage{
			Language: lang.Name,
			Repos:    r,
		})
	}

	if err = m.Stat.PostMostRecentReposByLanguage(ctx, repos); err != nil {
		return
	}
	return
}

// UpdateReposByLanguage updates the analytic type `repos_by_language`
func (m *model) UpdateReposByLanguage(ctx context.Context, perPage int) (err error) {
	var user []schema.UserCount
	var users []schema.UserCountByLanguage
	var languages []schema.LanguageCount

	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateReposByLanguage"),
				zap.Duration("took", time.Since(start)),
				zap.Int("perPage", perPage))

		if err != nil {
			zlog.Warn("error updating repos by language", zap.Error(err))
		} else {
			zlog.Info("update repos by language", zap.Int("count", len(users)))
		}
	}(time.Now())

	languages, err = m.Repo.MostPopularLanguage(ctx, perPage)
	if err != nil {
		return
	}

	for _, lang := range languages {
		user, err = m.Repo.ReposByLanguage(ctx, lang.Name, perPage)
		if err != nil {
			return
		}

		users = append(users, schema.UserCountByLanguage{
			Language: lang.Name,
			Users:    user,
		})
	}

	if err = m.Stat.PostReposByLanguage(ctx, users); err != nil {
		return
	}

	return
}

func (m *model) UpdateProfile(ctx context.Context, numWorkers int) (err error) {
	var logins []string
	var profiles []usersvc.User

	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateProfile"),
				zap.Duration("took", time.Since(start)),
				zap.Int("numWorkers", numWorkers))

		if err != nil {
			zlog.Warn("error updating profile", zap.Error(err))
		} else {
			zlog.Info("update profile",
				zap.Int("logins", len(logins)),
				zap.Int("profiles", len(profiles)))
		}
	}(time.Now())

	logins, err = m.Repo.DistinctLogin(ctx)
	if err != nil {
		return
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

	fanIn := func(ctx context.Context, numWorkers int, in <-chan string) <-chan usersvc.User {
		c := make(chan usersvc.User)

		var wg sync.WaitGroup
		wg.Add(numWorkers)

		multiplex := func(in <-chan string) {
			defer wg.Done()
			for i := range in {
				select {
				case <-ctx.Done():
					return
				case c <- m.Repo.GetProfile(ctx, i):
				}
			}
		}

		for i := 0; i < numWorkers; i++ {
			go multiplex(in)
		}

		go func() {
			defer close(c)
			wg.Wait()
		}()

		return c
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for p := range fanIn(ctx, numWorkers, toStream(ctx, logins...)) {
		profiles = append(profiles, p)
	}

	if err = m.User.BulkUpdate(ctx, profiles); err != nil {
		return
	}

	return
}

func (m *model) UpdateMatches(ctx context.Context) (err error) {
	defer func(start time.Time) {
		zlog := logger.Wrap(ctx, m.logger).
			With(zap.String("method", "UpdateMatches"),
				zap.Duration("took", time.Since(start)))

		if err != nil {
			zlog.Warn("error updating matches", zap.Error(err))
		} else {
			zlog.Info("update matches")
		}
	}(time.Now())
	maxMatches := 20
	var users []usersvc.User
	users, err = m.User.WithRepos(ctx, 0)

	for i := 0; i < len(users); i++ {
		p1 := users[i]
		var matches []schema.User

		for j := 0; j < len(users); j++ {
			p2 := users[j]
			if i == j {
				continue
			}

			matches = append(matches, schema.User{
				Login:     p2.Login,
				AvatarURL: p2.AvatarURL,
				Score:     recsys(p1, p2),
			})
		}
		sort.SliceStable(matches, func(i, j int) bool {
			return matches[i].Score > matches[j].Score
		})
		users[i].Profile.Matches = matches[:take(len(matches), maxMatches)]
	}

	if err = m.User.BulkUpdate(ctx, users); err != nil {
		return
	}

	return
}

func take(curr, max int) int {
	if curr < max {
		return curr
	}
	return max
}

func square(a, b int64) float64 {
	return math.Pow(float64(a-b), 2.0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func recsys(user1, user2 usersvc.User) (score float64) {
	sumSquares := square(user1.Repositories, user2.Repositories) +
		square(user1.Gists, user2.Gists) +
		square(user1.Followers, user2.Followers) +
		square(user1.Following, user2.Following) +
		square(user1.Profile.Watchers, user2.Profile.Watchers) +
		square(user1.Profile.Stargazers, user2.Profile.Stargazers) +
		square(user1.Profile.Forks, user2.Profile.Forks)

	mapAllLang := make(map[string]int64)
	mapLangUser1 := make(map[string]int64)
	mapLangUser2 := make(map[string]int64)

	for _, l := range user1.Languages {
		mapAllLang[l.Name]++
		mapLangUser1[l.Name] = int64(l.Count)
	}

	for _, l := range user2.Languages {
		mapAllLang[l.Name]++
		mapLangUser2[l.Name] = int64(l.Count)
	}

	for k, v := range mapAllLang {
		if v == 2 {
			sumSquares += square(mapLangUser1[k], mapLangUser2[k])
		}
	}

	mapKeywords := make(map[string]int64)
	mapKeywordUser1 := make(map[string]int64)
	mapKeywordUser2 := make(map[string]int64)

	for _, k := range user1.Keywords {
		mapKeywords[k.ID]++
		mapKeywordUser1[k.ID]++
	}

	for _, k := range user2.Languages {
		mapKeywords[k.Name]++
		mapKeywordUser2[k.Name]++
	}

	for k, v := range mapKeywords {
		if v == 2 {
			sumSquares += square(mapKeywordUser1[k], mapKeywordUser2[k])
		}
	}
	return 1 / (1 + math.Sqrt(sumSquares))
}
