package main

import (
	"context"
	stdlog "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/alextanhongpin/go-github-scraper/internal/app/mediatorsvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/reposvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/statsvc"
	"github.com/alextanhongpin/go-github-scraper/internal/app/transport"
	"github.com/alextanhongpin/go-github-scraper/internal/app/usersvc"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/client/github"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/cronjob"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/logger"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/null"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/profiler"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("version", "0.0.1")                             // The application version, normally the git hash
	viper.SetDefault("crontab_user_tab", "*/20 * * * * *")           // The crontab for user, running every 20 seconds
	viper.SetDefault("reset_repo", false)                            // Whether to fetch it from scratch or not
	viper.SetDefault("crontab_repo_tab", "0 * * * * *")              // The crontab for repo, running every minute
	viper.SetDefault("crontab_stat_tab", "0 10 0 * * *")             // The crontab for stat, running ten minutes after midnight
	viper.SetDefault("crontab_profile_tab", "@midnight")             // The crontab for profile, running at midnight
	viper.SetDefault("crontab_match_tab", "0 15 0 * * *")            // The crontab for matching, running fifteen minutes after midnight
	viper.SetDefault("crontab_user_enable", false)                   // The enable state of the crontab for user
	viper.SetDefault("crontab_repo_enable", false)                   // The enable state of the crontab for repo
	viper.SetDefault("crontab_stat_enable", false)                   // The enable state of the crontab for stat
	viper.SetDefault("crontab_profile_enable", false)                // The enable state of the crontab for profile
	viper.SetDefault("crontab_match_enable", false)                  // The enable state of the crontab for profile
	viper.SetDefault("crontab_user_trigger", false)                  // Will run once if set to true
	viper.SetDefault("crontab_repo_trigger", false)                  // Will run once if set to true
	viper.SetDefault("crontab_stat_trigger", false)                  // Will run once if set to true
	viper.SetDefault("crontab_profile_trigger", false)               // Will run once if set to true
	viper.SetDefault("crontab_match_trigger", false)                 // Will run once if set to true
	viper.SetDefault("db_user", "root")                              // The username of the database
	viper.SetDefault("db_pass", "example")                           // The password of the database
	viper.SetDefault("db_name", "scraper")                           // The name of the database
	viper.SetDefault("db_auth", "admin")                             // The name of the auth database
	viper.SetDefault("db_host", "mongodb://localhost:27017")         // The URI of the database
	viper.SetDefault("github_location", "Malaysia")                  // The default country to scrape data from
	viper.SetDefault("github_token", "")                             // The Github's access token used to make call to the GraphQL Endpoint
	viper.SetDefault("github_uri", "https://api.github.com/graphql") // The Github's GraphQL Endpoint
	viper.SetDefault("port", ":8080")                                // The TCP port of the application
	viper.SetDefault("pprof_port", ":6060")                          // The TCP port of for the http profiling
	viper.SetDefault("pprof_enable", false)                          // Toggle flag for pprof
	viper.SetDefault("cpuprofile", "")                               // Write cpuprofile to file, e.g. cpu.prof
	viper.SetDefault("memprofile", "")                               // Write memoryprofile to file, e.g. mem.prof
	viper.SetDefault("httpprofile", false)                           // Toggle state for http profiler
	viper.SetDefault("graceful_timeout", 15)                         // The duration for which the server gracefully wait for existing connections to finish
	viper.SetDefault("trace_endpoint", "http://localhost:14268")     // The endpoint of the jaeger image
	viper.SetDefault("trace_service", "go-scraper")                  // The name of the service that appears in the dashboard

	if viper.GetString("github_token") == "" {
		panic("github_token environment variable is missing")
	}
}

func main() {
	// Create global context for cancellation
	ctx := context.Background()

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

	// Setup jaeger
	exporter, err := jaeger.NewExporter(jaeger.Options{
		Endpoint:    viper.GetString("trace_endpoint"),
		ServiceName: viper.GetString("trace_service"),
	})

	if err != nil {
		stdlog.Fatal(err)
	}

	trace.RegisterExporter(exporter)

	// For demoing purpose, always sample
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	// Setup logger
	l := logger.New()
	defer l.Sync()

	// Setup database
	db := database.New(
		viper.GetString("db_host"),
		viper.GetString("db_user"),
		viper.GetString("db_pass"),
		viper.GetString("db_name"),
		viper.GetString("db_auth"))
	defer db.Close()

	// Setup services
	m := mediatorsvc.Mediator{
		Stat: statsvc.New(db,
			statsvc.Logging(l.Named("statsvc")),
			statsvc.Tracing()),
		Github: github.New(httpClient,
			viper.GetString("github_token"),
			viper.GetString("github_uri"),
			github.Logging(l.Named("github")),
			github.Tracing()),
		Repo: reposvc.New(db,
			reposvc.Logging(l.Named("reposvc")),
			reposvc.Tracing()),
		User: usersvc.New(db,
			usersvc.Logging(l.Named("usersvc")),
			usersvc.Tracing()),
	}

	// Setup mediator services, which is basically an orchestration of multiple services
	msvc := mediatorsvc.New(
		m,
		mediatorsvc.Logging(l.Named("mediatorsvc")),
		mediatorsvc.Tracing())

	// Setup cronjob
	cronjob.Exec(ctx,
		&cronjob.Config{
			Name:        "Fetch Users",
			Description: "Fetch the Github users data periodically based on location and created date, which is stored as delta timestamp",
			Start:       viper.GetBool("crontab_user_enable"),
			CronTab:     viper.GetString("crontab_user_tab"),
			Trigger:     viper.GetBool("crontab_user_trigger"),
			Fn: func(ctx context.Context) error {
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
			CronTab:     viper.GetString("crontab_repo_tab"),
			Trigger:     viper.GetBool("crontab_repo_trigger"),
			Fn: func(ctx context.Context) error {
				ctx = logger.WrapContextWithRequestID(ctx)
				userPerPage := 100
				repoPerPage := 30
				return msvc.FetchRepos(ctx, userPerPage, repoPerPage, viper.GetBool("reset_repo"))
			},
		},
		&cronjob.Config{
			Name:        "Update Profile",
			Description: "Compute the new user profile based on the repos that are scraped daily",
			Start:       viper.GetBool("crontab_profile_enable"),
			CronTab:     viper.GetString("crontab_profile_tab"),
			Trigger:     viper.GetBool("crontab_profile_trigger"),
			Fn: func(ctx context.Context) error {
				ctx = logger.WrapContextWithRequestID(ctx)
				numWorkers := 4
				return msvc.UpdateProfile(ctx, numWorkers)
			},
		},
		&cronjob.Config{
			Name:        "Build Stats",
			Description: "Compute the Github's analytic data of users in Malaysia based on the new repos that are scraped daily",
			Start:       viper.GetBool("crontab_stat_enable"),
			CronTab:     viper.GetString("crontab_stat_tab"),
			Trigger:     viper.GetBool("crontab_stat_trigger"),
			Fn: func(ctx context.Context) error {
				ctx = logger.WrapContextWithRequestID(ctx)
				defaultLimit := 20
				min := 3
				max := 100

				nullFns := []null.Fn{
					func() error { return msvc.UpdateUserCount(ctx) },
					func() error { return msvc.UpdateRepoCount(ctx) },
					func() error { return msvc.UpdateReposMostRecent(ctx, defaultLimit) },
					func() error { return msvc.UpdateRepoCountByUser(ctx, defaultLimit) },
					func() error { return msvc.UpdateReposMostStars(ctx, defaultLimit) },
					func() error { return msvc.UpdateReposMostForks(ctx, defaultLimit) },
					func() error { return msvc.UpdateLanguagesMostPopular(ctx, defaultLimit) },
					func() error { return msvc.UpdateMostRecentReposByLanguage(ctx, defaultLimit) },
					func() error { return msvc.UpdateReposByLanguage(ctx, defaultLimit) },
					func() error { return msvc.UpdateCompanyCount(ctx) },
					func() error { return msvc.UpdateUsersByCompany(ctx, min, max) },
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
		&cronjob.Config{
			Name:        "Update Matches",
			Description: "Compute the new user recommendations based on the new repos pulled",
			Start:       viper.GetBool("crontab_match_enable"),
			CronTab:     viper.GetString("crontab_match_tab"),
			Trigger:     viper.GetBool("crontab_match_trigger"),
			Fn: func(ctx context.Context) error {
				ctx = logger.WrapContextWithRequestID(ctx)
				return msvc.UpdateMatches(ctx)
			},
		},
	)

	// Setup router
	r := httprouter.New()

	// Setup endpoints, can also add feature toggle capabilities
	tr := transport.New(r)
	tr.Init(
		transport.NewUserEndpoints(m.User),
		transport.NewStatEndpoints(m.Stat),
		transport.NewRepoEndpoints(m.Repo),
	)

	// Add cors support
	handler := cors.Default().Handler(r)

	// a http.Server with pre-configured timeouts to avoid Slowloris attack
	srv := &http.Server{
		Addr:           viper.GetString("port"),
		Handler:        handler,
		ReadTimeout:    time.Second * 10, // Variable always on the right, not 10 * time.Second
		WriteTimeout:   time.Second * 10,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}

	// Setup pprof net/http
	if viper.GetBool("pprof_enable") {
		go func() {
			stdlog.Fatal(http.ListenAndServe(viper.GetString("pprof_port"), nil))
		}()
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
	ctx, cancel := context.WithTimeout(ctx, time.Second*viper.GetDuration("graceful_timeout"))
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout
	srv.Shutdown(ctx)

	stdlog.Println("shutting down server")
	os.Exit(0)
}
