package github

import (
	"log"
)

// Model represents the api interface for the Github's GraphQL
type model struct {
	store Store
}

// NewModel returns a new model
func NewModel(store Store) API {
	return &model{
		store: store,
	}
}

func (m *model) FetchUsersCursor(location, start, end string, limit int) ([]User, error) {
	log.Printf(`api.github.model.FetchUsersCursor location:%s start:%s end:%s limit:%d\n`, location, start, end, limit)
	cursor := ""
	hasNextPage := true
	var users []User
	for hasNextPage {
		res, err := m.store.FetchUsers(FetchUsersRequest{
			Location: location,
			Start:    start,
			End:      end,
			Cursor:   cursor,
			Limit:    limit,
		})
		if err != nil {
			log.Println(err)
			break
		}
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		userCount := res.Data.Search.UserCount

		log.Printf("hasNextPage:%v userCount:%d\n", hasNextPage, userCount)
		for _, edge := range res.Data.Search.Edges {
			users = append(users, edge.Node)
		}
		if hasNextPage {
			log.Printf("fetching next cursor: %s\n", cursor)
		}
	}
	log.Printf("api.github.model.FetchUsersCursor count:%d\n", len(users))
	return users, nil
}

func (m *model) FetchReposCursor(login, start, end string, limit int) ([]Repo, error) {
	cursor := ""
	hasNextPage := true
	var repos []Repo
	for hasNextPage {
		res, err := m.store.FetchRepos(FetchReposRequest{
			Login:  login,
			Start:  start,
			End:    end,
			Cursor: cursor,
			Limit:  limit,
		})
		if err != nil {
			log.Println("api.github.model.FetchReposCursor error:", err)
			break
		}
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		repoCount := res.Data.Search.RepositoryCount
		log.Printf("api.github.model.FetchReposCursor hasNextPage:%v repoCount:%d\n", hasNextPage, repoCount)
		for _, edge := range res.Data.Search.Edges {
			repos = append(repos, edge.Node)
		}
	}
	log.Printf("api.github.model.FetchReposCursor totalCount:%d\n", len(repos))
	return repos, nil
}
