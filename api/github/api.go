package github

import "log"

// API represents the api interface for the Github's GraphQL
type API interface {
	FetchUsersCursor(location, start, end string) ([]User, error)
	FetchReposCursor(login, start, end string) ([]Repo, error)
}

type api struct {
	store Store
}

// NewAPI returns a new api
func NewAPI(store Store) API {
	return &api{
		store: store,
	}
}

func (a *api) FetchUsersCursor(location, start, end string) ([]github.User, error) {
	cursor := ""
	hasNextPage := true
	var users []github.User
	for hasNextPage {
		res, err := a.store.FetchUsers(location, start, end, cursor)
		if err != nil {
			log.Println(err)
			break
		}
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		for _, edge := range res.Data.Search.Edges {
			users = append(users, edge.Node)
		}
	}
	return users, nil
}

func (a *api) FetchReposCursor(login, start, end string) ([]github.Repo, error) {
	cursor := ""
	hasNextPage := true
	var repos []github.Repo
	for hasNextPage {
		res, err := a.store.FetchRepos(login, start, end, cursor)
		if err != nil {
			log.Println(err)
			break
		}
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		for _, edge := range res.Data.Search.Edges {
			repos = append(repos, edge.Node)
		}
	}
	return repos, nil
}
