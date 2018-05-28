package main

import (
	"context"
	stdlog "log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/app/cronjob"
	"github.com/alextanhongpin/go-github-scraper/internal/app/mediator"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/database"
	"github.com/alextanhongpin/go-github-scraper/internal/pkg/profiler"
	"github.com/alextanhongpin/go-github-scraper/internal/util"
	"github.com/alextanhongpin/go-github-scraper/service/analyticsvc"
	"github.com/alextanhongpin/go-github-scraper/service/profilesvc"
	"github.com/alextanhongpin/go-github-scraper/service/reposvc"
	"github.com/alextanhongpin/go-github-scraper/service/usersvc"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("crontab_user", "*/20 * * * * *")                     // The crontab for user, running every 20 seconds
	viper.SetDefault("crontab_repo", "0 * * * * *")                        // The crontab for repo, running every minute
	viper.SetDefault("crontab_analytic", "@daily")                         // The crontab for analytic, running daily
	viper.SetDefault("crontab_user_enable", false)                         // The enable state of the crontab for user
	viper.SetDefault("crontab_repo_enable", false)                         // The enable state of the crontab for repo
	viper.SetDefault("crontab_analytic_enable", false)                     // The enable state of the crontab for analytic
	viper.SetDefault("db_name", "scraper")                                 // The name of the database
	viper.SetDefault("db_host", "mongodb://myuser:mypass@localhost:27017") // The URI of the database
	viper.SetDefault("github_created_at", "2008-04-01")                    // Github's created date, used as default date for scraping
	viper.SetDefault("github_location", "Malaysia")                        // The default country to scrape data from
	viper.SetDefault("github_token", "")                                   // The Github's access token used to make call to the GraphQL Endpoint
	viper.SetDefault("github_uri", "https://api.github.com/graphql")       // The Github's GraphQL Endpoint
	viper.SetDefault("port", ":8080")                                      // The TCP port of the application
	viper.SetDefault("cpuprofile", "cpu.prof")                             // Write cpuprofile to file, e.g. cpu.prof
	viper.SetDefault("memprofile", "mem.prof")                             // Write memoryprofile to file, e.g. mem.prof
	viper.SetDefault("graceful_timeout", "15")                             // The duration for which the server gracefully wait for existing connections to finish
	if viper.GetString("github_token") == "" {
		panic("github_token environment variable is missing")
	}
}

func main() {
	profiler.MakeCPU(viper.GetString("cpuprofile"))

	// Setup Logger
	logger, err := zap.NewProduction()
	if err != nil {
		stdlog.Fatal(err)
	}
	defer logger.Sync()

	// Setup Database
	db := database.New(viper.GetString("db_host"), viper.GetString("db_name"))
	defer db.Close()

	// Setup services
	m := mediator.Mediator{
		Analytic: analyticsvc.New(db),
		Github: github.New(util.NewHTTPClient(),
			viper.GetString("github_token"),
			viper.GetString("github_uri"),
			logger),
		Profile: profilesvc.New(db),
		Repo:    reposvc.New(db, logger),
		User:    usersvc.New(db),
	}

	// Setup mediator services, which is basically an orchestration of multiple services
	msvc := mediator.New(m, logger)

	cronjob.Exec(
		&cronjob.Config{
			Name:        "Fetch Users",
			Description: "Fetch the Github users data periodically based on location and created date, which is stored as delta timestamp",
			Start:       false,
			CronTab:     "* * * * * *",
			Fn: func() error {
				months := 6
				perPage := 30
				return msvc.FetchUsers("Malaysia", months, perPage)
			},
		},
		&cronjob.Config{
			Name:        "Fetch Repos",
			Description: "Fetch the Github user's repos periodically based on the last fetched date",
			Start:       false,
			CronTab:     "* * * * * *",
			Fn: func() error {
				userPerPage := 30
				repoPerPage := 30
				return msvc.FetchRepos(userPerPage, repoPerPage)
			},
		},
		&cronjob.Config{
			Name:        "Update Profile",
			Description: "Compute the new user profile based on the repos that are scraped daily",
			Start:       false,
			CronTab:     "* * * * * *",
			Fn: func() error {
				numWorkers := 16
				return msvc.UpdateProfile(numWorkers)
			},
		},
		&cronjob.Config{
			Name:        "Build Analytic",
			Description: "Compute the Github's analytic data of users in Malaysia based on the new repos that are scraped daily",
			Start:       viper.GetBool("crontab_analytic_enable"),
			CronTab:     viper.GetString("crontab_analytic"),
			Fn: func() error {
				var wg sync.WaitGroup
				wg.Add(8)
				go func() {
					msvc.UpdateUserCount()
					wg.Done()
				}()
				go func() {
					msvc.UpdateRepoCount()
					wg.Done()
				}()
				go func() {
					msvc.UpdateReposMostRecent(20)
					wg.Done()
				}()
				go func() {
					msvc.UpdateRepoCountByUser(20)
					wg.Done()
				}()
				go func() {
					msvc.UpdateReposMostStars(20)
					wg.Done()
				}()
				go func() {
					msvc.UpdateLanguagesMostPopular(20)
					wg.Done()
				}()
				go func() {
					msvc.UpdateMostRecentReposByLanguage(20)
					wg.Done()
				}()
				go func() {
					msvc.UpdateReposByLanguage(20)
					wg.Done()
				}()
				wg.Wait()
				return nil
			},
		},
	)

	// Setup router
	r := httprouter.New()

	// Setup profiling
	isProfilerEnabled := true
	profiler.MakeHTTP(isProfilerEnabled, r)

	// Setup endpoints, can also add feature toggle capabilities
	usersvc.MakeEndpoints(m.User, r) // A better way? - usvc.Wrap(r), usersvc.Bind(usvc, r)
	analyticsvc.MakeEndpoints(m.Analytic, r)
	reposvc.MakeEndpoints(m.Repo, r)

	// Setup server
	srv := util.NewHTTPServer(viper.GetString("port"), r)
	// Run our server in a goroutine so that it doesn't block

	go func() {
		stdlog.Printf("listening to port *%s. press ctrl + c to cancel.\n", viper.GetString("port"))
		stdlog.Fatal(srv.ListenAndServe())
	}()

	profiler.MakeMemory(viper.GetString("memprofile"))

	c := make(chan os.Signal, 1)

	// Accept graceful shutdowns when quit via SIGINT (Ctrl + C) SIGKILL,
	// SIGQUIT or SIGTERM (Ctrl + /) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("graceful_timeout")*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout
	srv.Shutdown(ctx)

	stdlog.Println("shutting down")
	os.Exit(0)
}
