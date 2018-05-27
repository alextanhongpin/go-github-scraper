package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"
	rpprof "runtime/pprof"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/internal/database"
	"github.com/alextanhongpin/go-github-scraper/internal/util"
	"github.com/alextanhongpin/go-github-scraper/service/analyticsvc"
	"github.com/alextanhongpin/go-github-scraper/service/profilesvc"
	"github.com/alextanhongpin/go-github-scraper/service/reposvc"
	"github.com/alextanhongpin/go-github-scraper/service/usersvc"
	"github.com/alextanhongpin/go-github-scraper/worker"

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
	if viper.GetString("github_token") == "" {
		panic("github_token environment variable is missing")
	}
}

func main() {
	if viper.GetString("cpuprofile") != "" {
		f, err := os.Create(viper.GetString("cpuprofile"))
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := rpprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer rpprof.StopCPUProfile()
	}

	// Setup Logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	// Setup Database
	db := database.New(viper.GetString("db_host"), viper.GetString("db_name"))
	defer db.Close()

	// Setup services
	asvc := analyticsvc.New(db)
	rsvc := reposvc.New(db)
	usvc := usersvc.New(db)
	psvc := profilesvc.New(db)
	gsvc := github.New(
		viper.GetString("github_token"),
		viper.GetString("github_uri"))

	// Setup workers
	w := worker.New(gsvc, asvc, psvc, rsvc, usvc, logger)

	logger.Info("crontab_user", zap.Bool("is_enabled", viper.GetBool("crontab_user_enable")))
	if viper.GetBool("crontab_user_enable") {
		c1 := w.NewFetchUsers(viper.GetString("crontab_user"))
		c1.Start()
	}

	logger.Info("crontab_repo", zap.Bool("is_enabled", viper.GetBool("crontab_repo_enable")))
	if viper.GetBool("crontab_repo_enable") {
		c2 := w.NewFetchRepos(viper.GetString("crontab_repo"))
		c2.Start()
	}

	logger.Info("crontab_analytic", zap.Bool("is_enabled", viper.GetBool("crontab_analytic_enable")))
	if viper.GetBool("crontab_analytic_enable") {
		c3 := w.NewAnalyticBuilder(viper.GetString("crontab_analytic"))
		c3.Start()
	}

	_ = w.NewProfileBuilder(viper.GetString("crontab_analytic"))
	// c4.Start()

	// Setup router
	r := httprouter.New()

	// Setup profiling
	r.HandlerFunc(http.MethodGet, "/debug/pprof/", pprof.Index)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/cmdline", pprof.Cmdline)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/profile", pprof.Profile)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/symbol", pprof.Symbol)
	r.HandlerFunc(http.MethodGet, "/debug/pprof/trace", pprof.Trace)
	r.Handler(http.MethodGet, "/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handler(http.MethodGet, "/debug/pprof/heap", pprof.Handler("heap"))
	r.Handler(http.MethodGet, "/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handler(http.MethodGet, "/debug/pprof/block", pprof.Handler("block"))

	// Setup endpoints, can also add feature toggle capabilities
	usersvc.MakeEndpoints(usvc, r)
	analyticsvc.MakeEndpoints(asvc, r)
	reposvc.MakeEndpoints(rsvc, r)

	if viper.GetString("memprofile") != "" {
		f, err := os.Create(viper.GetString("memprofile"))
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := rpprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	// Setup server
	server := util.NewHTTPServer(viper.GetString("port"), r)
	log.Printf("listening to port *%s. press ctrl + c to cancel.\n", viper.GetString("port"))
	log.Fatal(server.ListenAndServe())
}
