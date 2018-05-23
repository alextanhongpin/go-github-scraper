package main

import (
	"flag"
	"log"

	"github.com/alextanhongpin/go-github-scraper/api/github"
	"github.com/alextanhongpin/go-github-scraper/model"
)

func main() {
	githubToken := flag.String("github_token", "", "The Github's access token used to make calls to the GraphQL endpoint")
	flag.Parse()

	gapi := github.New(*githubToken, "https://api.github.com/graphql", "Malaysia")
	cursor := ""
	hasNextPage := true
	var users []model.User
	for hasNextPage {
		res, err := gapi.FetchUsers("2018-01-01", "2018-02-01", cursor)
		if err != nil {
			break
		}
		log.Println("got user count:", res.Data.Search.UserCount)
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		for _, edge := range res.Data.Search.Edges {
			users = append(users, edge.Node)
		}
	}
	log.Println(len(users))
}
