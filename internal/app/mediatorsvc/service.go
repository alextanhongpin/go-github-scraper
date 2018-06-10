// Package mediatorsvc handles multiple service orchestrations
// and is categorized as domain model
package mediatorsvc

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/app/reposvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/statsvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/heapsort"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/moment"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/schema"
)

type (
	// Service represents the methods the mediator service must implement
	Service interface {
		FetchUsers(ctx context.Context, location string, months int, perPage int) error
		FetchRepos(ctx context.Context, userPerPage, repoPerPage int, reset bool) error
		UpdateUserCount(ctx context.Context) error
		UpdateRepoCount(ctx context.Context) error
		UpdateReposMostRecent(ctx context.Context, perPage int) error
		UpdateRepoCountByUser(ctx context.Context, perPage int) error
		UpdateReposMostStars(ctx context.Context, perPage int) error
		UpdateReposMostForks(ctx context.Context, perPage int) error
		UpdateLanguagesMostPopular(ctx context.Context, perPage int) error
		UpdateMostRecentReposByLanguage(ctx context.Context, perPage int) error
		UpdateReposByLanguage(ctx context.Context, perPage int) error
		UpdateProfile(ctx context.Context, numWorkers int) error
		UpdateMatches(ctx context.Context) error
		UpdateUsersByCompany(ctx context.Context, min, max int) error
		UpdateCompanyCount(ctx context.Context) error
	}

	// Mediator holds the services in used
	Mediator struct {
		Github github.Service
		Stat   statsvc.Service
		Repo   reposvc.Service
		User   usersvc.Service
	}

	service struct {
		Mediator
	}
)

func NewService(m Mediator) Service {
	return &service{m}
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

func (s *service) FetchUsers(ctx context.Context, location string, months int, perPage int) error {
	start, _ := s.User.FindLastCreated(ctx)
	end := makeEndDate(start, months)

	users, err := s.Github.FetchUsersCursor(ctx, location, start, end, perPage)
	if err != nil {
		return err
	}

	return s.User.BulkUpsert(ctx, users)
}

func (s *service) FetchRepos(ctx context.Context, userPerPage, repoPerPage int, reset bool) error {
	users, err := s.User.FindLastFetched(ctx, userPerPage)
	if err != nil {
		return err
	}

	for _, user := range users {
		login := user.Login
		if login == "" {
			continue
		}

		start, _ := s.Repo.LastCreatedBy(ctx, login)
		end := moment.NewCurrentFormattedDate()

		repos, err := s.Github.FetchReposCursor(ctx, login, start, end, repoPerPage)
		if err != nil {
			return err
		}

		if err = s.Repo.BulkUpsert(ctx, repos); err != nil {
			return err
		}

		if err = s.User.UpdateOne(ctx, login); err != nil {
			return err
		}
	}
	return nil
}

// UpdateUserCount updates the analytic type `user_count`
func (s *service) UpdateUserCount(ctx context.Context) error {
	count, err := s.User.Count(ctx)
	if err != nil {
		return err
	}

	return s.Stat.PostUserCount(ctx, count)
}

// UpdateRepoCount updates the analytic type `repo_count`
func (s *service) UpdateRepoCount(ctx context.Context) error {
	count, err := s.Repo.Count(ctx)
	if err != nil {
		return err
	}

	return s.Stat.PostRepoCount(ctx, count)
}

// UpdateReposMostRecent updates the analytic type `repos_most_recent`
func (s *service) UpdateReposMostRecent(ctx context.Context, perPage int) error {
	repos, err := s.Repo.MostRecent(ctx, perPage)
	if err != nil {
		return err
	}

	return s.Stat.PostReposMostRecent(ctx, repos)
}

// UpdateRepoCountByUser updates the analytic type `repo_count_by_user`
func (s *service) UpdateRepoCountByUser(ctx context.Context, perPage int) error {
	users, err := s.Repo.RepoCountByUser(ctx, perPage)
	if err != nil {
		return err
	}

	return s.Stat.PostRepoCountByUser(ctx, users)
}

// UpdateReposMostStars updates the analytic type `repos_most_stars`
func (s *service) UpdateReposMostStars(ctx context.Context, perPage int) error {
	repos, err := s.Repo.MostStars(ctx, perPage)
	if err != nil {
		return err
	}

	return s.Stat.PostReposMostStars(ctx, repos)
}

// UpdateReposMostForks updates the analytic type `repos_most_forks`
func (s *service) UpdateReposMostForks(ctx context.Context, perPage int) error {
	repos, err := s.Repo.MostForks(ctx, perPage)
	if err != nil {
		return err
	}

	return s.Stat.PostReposMostForks(ctx, repos)
}

// UpdateLanguagesMostPopular updates the analytic type `languages_most_popular`
func (s *service) UpdateLanguagesMostPopular(ctx context.Context, perPage int) error {
	languages, err := s.Repo.MostPopularLanguage(ctx, perPage)
	if err != nil {
		return err
	}

	return s.Stat.PostMostPopularLanguage(ctx, languages)
}

// UpdateMostRecentReposByLanguage updates the analytic type `repos_most_recent_by_language`
func (s *service) UpdateMostRecentReposByLanguage(ctx context.Context, perPage int) error {
	languages, err := s.Repo.MostPopularLanguage(ctx, perPage)
	if err != nil {
		return err
	}

	var repos []schema.RepoLanguage
	for _, lang := range languages {
		r, err := s.Repo.MostRecentReposByLanguage(ctx, lang.Name, perPage)
		if err != nil {
			return err
		}
		repos = append(repos, schema.RepoLanguage{
			Language: lang.Name,
			Repos:    r,
		})
	}

	return s.Stat.PostMostRecentReposByLanguage(ctx, repos)
}

// UpdateReposByLanguage updates the analytic type `repos_by_language`
func (s *service) UpdateReposByLanguage(ctx context.Context, perPage int) error {
	languages, err := s.Repo.MostPopularLanguage(ctx, perPage)
	if err != nil {
		return err
	}

	var users []schema.UserCountByLanguage
	for _, lang := range languages {
		user, err := s.Repo.ReposByLanguage(ctx, lang.Name, perPage)
		if err != nil {
			return err
		}

		users = append(users, schema.UserCountByLanguage{
			Language: lang.Name,
			Users:    user,
		})
	}

	return s.Stat.PostReposByLanguage(ctx, users)
}

func (s *service) UpdateProfile(ctx context.Context, numWorkers int) error {
	logins, err := s.Repo.Distinct(ctx, "login")
	if err != nil {
		return err
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
				p, err := s.Repo.GetProfile(ctx, i)
				if err != nil {
					continue
				}
				select {
				case <-ctx.Done():
					return
				case c <- *p:
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

	var profiles []usersvc.User
	for p := range fanIn(ctx, numWorkers, toStream(ctx, logins...)) {
		profiles = append(profiles, p)
	}

	return s.User.BulkUpdate(ctx, profiles)
}

func (s *service) UpdateMatches(ctx context.Context) error {

	maxMatches := 20
	var users []usersvc.User
	users, err := s.User.WithRepos(ctx, 0)
	if err != nil {
		return err
	}

	for i := 0; i < len(users); i++ {
		p1 := users[i]
		var matches []schema.User

		for j := 0; j < len(users); j++ {
			p2 := users[j]
			if i == j {
				continue
			}

			score := recsys(p1, p2)

			if len(matches) > maxMatches {
				if score > matches[0].Score {
					matches = append(matches[1:], schema.User{
						Login:     p2.Login,
						AvatarURL: p2.AvatarURL,
						Score:     score,
					})
					heapsort.Sort(matches)
				}
			} else {
				matches = append(matches, schema.User{
					Login:     p2.Login,
					AvatarURL: p2.AvatarURL,
					Score:     score,
				})
			}
		}
		users[i].Profile.Matches = matches[:take(len(matches), maxMatches)]
	}

	return s.User.BulkUpdate(ctx, users)
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

func (s *service) UpdateCompanyCount(ctx context.Context) error {
	var res []string
	res, err := s.User.DistinctCompany(ctx)
	if err != nil {
		return err
	}

	return s.Stat.PostCompanyCount(ctx, len(res))
}

func (s *service) UpdateUsersByCompany(ctx context.Context, min, max int) error {
	companies, err := s.User.AggregateCompany(ctx, min, max)
	if err != nil {
		return err
	}

	for i := 0; i < len(companies); i++ {
		res, err := s.User.FindByCompany(ctx, companies[i].Company)
		if err != nil {
			continue
		}
		companies[i].Users = res
	}

	return s.Stat.PostUsersByCompany(ctx, companies)
}
