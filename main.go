package main

import (
	"context"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/app/mediatorsvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/profilesvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/reposvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/statsvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/cronjob"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/null"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/profiler"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("crontab_user", "*/20 * * * * *")                     // The crontab for user, running every 20 seconds
	viper.SetDefault("crontab_repo", "0 * * * * *")                        // The crontab for repo, running every minute
	viper.SetDefault("crontab_stat", "@daily")                             // The crontab for stat, running daily
	viper.SetDefault("crontab_profile", "@daily")                          // The crontab for profile, running daily
	viper.SetDefault("crontab_user_enable", false)                         // The enable state of the crontab for user
	viper.SetDefault("crontab_repo_enable", false)                         // The enable state of the crontab for repo
	viper.SetDefault("crontab_stat_enable", false)                         // The enable state of the crontab for stat
	viper.SetDefault("crontab_profile_enable", false)                      // The enable state of the crontab for profile
	viper.SetDefault("db_name", "scraper")                                 // The name of the database
	viper.SetDefault("db_host", "mongodb://myuser:mypass@localhost:27017") // The URI of the database
	viper.SetDefault("github_location", "Malaysia")                        // The default country to scrape data from
	viper.SetDefault("github_token", "")                                   // The Github's access token used to make call to the GraphQL Endpoint
	viper.SetDefault("github_uri", "https://api.github.com/graphql")       // The Github's GraphQL Endpoint
	viper.SetDefault("port", ":8080")                                      // The TCP port of the application
	viper.SetDefault("cpuprofile", "cpu.prof")                             // Write cpuprofile to file, e.g. cpu.prof
	viper.SetDefault("memprofile", "mem.prof")                             // Write memoryprofile to file, e.g. mem.prof
	viper.SetDefault("httpprofile", false)                                 // Toggle state for http profiler
	viper.SetDefault("graceful_timeout", "15")                             // The duration for which the server gracefully wait for existing connections to finish
	if viper.GetString("github_token") == "" {
		panic("github_token environment variable is missing")
	}
}

func main() {
	// Setup cpu profiler
	profiler.MakeCPU(viper.GetString("cpuprofile"))

	// Setup http client
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    20,
			IdleConnTimeout: time.Second * 5,
		},
		Timeout: time.Second * 5,
	}

	// Setup logger
	l := logger.New()
	defer l.Sync()

	// Setup database
	db := database.New(viper.GetString("db_host"), viper.GetString("db_name"))
	defer db.Close()

	// Setup services
	m := mediatorsvc.Mediator{
		Analytic: statsvc.New(db),
		Github: github.New(httpClient,
			viper.GetString("github_token"),
			viper.GetString("github_uri"),
			l),
		Profile: profilesvc.New(db),
		Repo:    reposvc.New(db),
		User:    usersvc.New(db),
	}

	// Setup mediator services, which is basically an orchestration of multiple services
	msvc := mediatorsvc.New(m)

	// Setup cronjob
	cronjob.Exec(
		&cronjob.Config{
			Name:        "Fetch Users",
			Description: "Fetch the Github users data periodically based on location and created date, which is stored as delta timestamp",
			Start:       viper.GetBool("crontab_user_enable"),
			CronTab:     viper.GetString("crontab_user"),
			Fn: func() error {
				ctx := context.Background()
				ctx = logger.WrapContextWithRequestID(ctx)
				location := "Malaysia"
				months := 6
				perPage := 30
				return msvc.FetchUsers(ctx, location, months, perPage)
			},
		},
		&cronjob.Config{
			Name:        "Fetch Repos",
			Description: "Fetch the Github user's repos periodically based on the last fetched date",
			Start:       viper.GetBool("crontab_repo_enable"),
			CronTab:     viper.GetString("crontab_repo"),
			Fn: func() error {
				ctx := context.Background()
				ctx = logger.WrapContextWithRequestID(ctx)
				userPerPage := 30
				repoPerPage := 30
				return msvc.FetchRepos(ctx, userPerPage, repoPerPage)
			},
		},
		&cronjob.Config{
			Name:        "Update Profile",
			Description: "Compute the new user profile based on the repos that are scraped daily",
			Start:       viper.GetBool("crontab_profile_enable"),
			CronTab:     viper.GetString("crontab_profile"),
			Fn: func() error {
				ctx := context.Background()
				ctx = logger.WrapContextWithRequestID(ctx)
				numWorkers := 16
				return msvc.UpdateProfile(ctx, numWorkers)
			},
		},
		&cronjob.Config{
			Name:        "Build Stats",
			Description: "Compute the Github's analytic data of users in Malaysia based on the new repos that are scraped daily",
			Start:       viper.GetBool("crontab_stat_enable"),
			CronTab:     viper.GetString("crontab_stat"),
			Fn: func() error {
				ctx := context.Background()
				ctx = logger.WrapContextWithRequestID(ctx)
				nullFns := []null.Fn{
					func() error { return msvc.UpdateUserCount(ctx) },
					func() error { return msvc.UpdateRepoCount(ctx) },
					func() error { return msvc.UpdateReposMostRecent(ctx, 20) },
					func() error { return msvc.UpdateRepoCountByUser(ctx, 20) },
					func() error { return msvc.UpdateReposMostStars(ctx, 20) },
					func() error { return msvc.UpdateLanguagesMostPopular(ctx, 20) },
					func() error { return msvc.UpdateMostRecentReposByLanguage(ctx, 20) },
					func() error { return msvc.UpdateReposByLanguage(ctx, 20) },
				}
				var wg sync.WaitGroup
				wg.Add(len(nullFns))

				for _, fn := range nullFns {
					go func(f null.Fn) {
						defer wg.Done()
						f()
					}(fn)
				}

				wg.Wait()
				return nil
			},
		},
	)

	// Setup router
	r := httprouter.New()

	// Setup http profiling
	profiler.MakeHTTP(viper.GetBool("httpprofile"), r)

	// Setup endpoints, can also add feature toggle capabilities
	usersvc.MakeEndpoints(m.User, r) // A better way? - usvc.Wrap(r), usersvc.Bind(usvc, r)
	statsvc.MakeEndpoints(m.Analytic, r)
	reposvc.MakeEndpoints(m.Repo, r)

	// a http.Server with pre-configured timeouts to avoid Slowloris attack
	srv := &http.Server{
		Addr:           viper.GetString("port"),
		Handler:        r,
		ReadTimeout:    time.Second * 10, // Variable always on the right, not 10 * time.Second
		WriteTimeout:   time.Second * 10,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}

	// Run our server in a goroutine so that it doesn't block
	go func() {
		stdlog.Printf("listening to port *%s. press ctrl + c to cancel.\n", viper.GetString("port"))
		stdlog.Fatal(srv.ListenAndServe())
	}()

	// Setup memory profiler
	profiler.MakeMemory(viper.GetString("memprofile"))

	c := make(chan os.Signal, 1)

	// Accept graceful shutdowns when quit via SIGINT (Ctrl + C) SIGKILL,
	// SIGQUIT or SIGTERM (Ctrl + /) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*viper.GetDuration("graceful_timeout"))
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout
	srv.Shutdown(ctx)

	stdlog.Println("shutting down server")
	os.Exit(0)
}
