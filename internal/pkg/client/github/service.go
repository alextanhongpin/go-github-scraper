package github

import (
	"context"
)

type (
	// Service represents the Github Service
	Service interface {
		FetchUsersCursor(ctx context.Context, location, start, end string, limit int) (users []User, err error)
		FetchReposCursor(ctx context.Context, login, start, end string, limit int) (repos []Repo, err error)
	}

	service struct {
		model Model
	}
)

// NewService returns a new service
func NewService(model Model) Service {
	return &service{model}
}

func (s *service) FetchUsersCursor(ctx context.Context, location, start, end string, limit int) ([]User, error) {
	var users []User

	cursor := ""
	hasNextPage := true
	for hasNextPage {
		res, err := s.model.FetchUsers(FetchUsersRequest{
			Location: location,
			Start:    start,
			End:      end,
			Cursor:   cursor,
			Limit:    limit,
		})
		if err != nil {
			return nil, err
		}
		hasNextPage = res.Data.Search.PageInfo.HasNextPage
		cursor = res.Data.Search.PageInfo.EndCursor
		for _, edge := range res.Data.Search.Edges {
			users = append(users, edge.Node)
		}
	}
	return users, nil
}

func (s *service) FetchReposCursor(ctx context.Context, login, start, end string, limit int) ([]Repo, error) {
	var repos []Repo
	cursor := ""
	hasNextPage := true

	for hasNextPage {
		res, err := s.model.FetchRepos(FetchReposRequest{
			Login:  login,
			Start:  start,
			End:    end,
			Cursor: cursor,
			Limit:  limit,
		})
		if err != nil {
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
