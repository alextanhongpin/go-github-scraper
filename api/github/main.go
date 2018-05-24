package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/model"
	"github.com/alextanhongpin/go-github-scraper/util"
)

// Service represents the interface for the Github Service
type Service interface {
	FetchUsers(location, start, end, cursor string) (*FetchUsersResponse, error)
	FetchRepos(login, start, end, cursor string) (*FetchReposResponse, error)
	FetchUsersCursor(location, start, end string) ([]model.User, error)
	FetchReposCursor(login, start, end string) ([]model.Repo, error)
}

// Service holds the service configuration
type service struct {
	client   *http.Client
	token    string
	endpoint string
}

// New returns a new service
func New(token, endpoint string) Service {
	return &service{
		client:   util.NewHTTPClient(),
		token:    token,
		endpoint: endpoint,
	}
}

func (s *service) FetchUsers(location, start, end, cursor string) (*FetchUsersResponse, error) {
	body := newUserQuery(location, start, end, cursor, 10)
	jsonBytes, err := json.Marshal(model.GraphQLQuery{body})
	if err != nil {
		log.Println(err)
	}

	jsonResp, err := graphqlService(s.client, s.token, s.endpoint, jsonBytes)
	if err != nil {
		return nil, err
	}
	var resp FetchUsersResponse
	if err := json.Unmarshal(jsonResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *service) FetchRepos(login, start, end, cursor string) (*FetchReposResponse, error) {
	body := newRepoQuery(login, start, end, cursor, 10)
	jsonBytes, err := json.Marshal(model.GraphQLQuery{body})
	if err != nil {
		log.Println(err)
	}

	jsonResp, err := graphqlService(s.client, s.token, s.endpoint, jsonBytes)
	if err != nil {
		return nil, err
	}
	log.Println(string(jsonResp))
	var resp FetchReposResponse
	if err := json.Unmarshal(jsonResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *service) FetchUsersCursor(location, start, end string) ([]model.User, error) {
	cursor := ""
	hasNextPage := true
	var users []model.User
	for hasNextPage {
		res, err := s.FetchUsers(location, start, end, cursor)
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

func (s *service) FetchReposCursor(login, start, end string) ([]model.Repo, error) {
	cursor := ""
	hasNextPage := true
	var repos []model.Repo
	for hasNextPage {
		res, err := s.FetchRepos(login, start, end, cursor)
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

func graphqlService(client *http.Client, token, endpoint string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func newUserQuery(location, start, end, cursor string, limit int) string {
	return fmt.Sprintf(`
		query {
			search(query: "location:%s created:%s..%s", type: USER, last: %d, after: %s) {
				userCount,
				pageInfo {
					startCursor,
					endCursor,
					hasNextPage,
					hasPreviousPage,
				},
				edges {
					cursor,
					node {
						...on User {
							name,
							createdAt,
							updatedAt,
							login,
							bio,
							location,
							email,
							company,
							avatarUrl,
							websiteUrl
						}
					}
				}
			}
		}`,
		location, start, end, limit, cursor)
}

func newRepoQuery(login, start, end, cursor string, limit int) string {
	return fmt.Sprintf(`query {
		search(query: "user:%s created:%s..%s", type: REPOSITORY, last: %d, after: %s) {
			repositoryCount,
			pageInfo {
				hasNextPage,
				startCursor,
				endCursor,
				hasPreviousPage,
			},
			edges {
				cursor,
				node {
					...on Repository {
						name,
						createdAt,
						updatedAt,
						description,
						homepageUrl,
						forkCount,
						isFork,
						nameWithOwner,
						languages (last: 30) {
							totalCount,
							edges {
								node {
									name,
									color
								}
							}
						},
						owner {
							login,
							avatarUrl
						},
						stargazers (last: 0) {
							totalCount
						},
						watchers (last: 0) {
							totalCount
						}
					}
				}
			}
		}
	}`, login, start, end, limit, cursor)
}

// FetchUsersResponse represents the GraphQL's user data structure
type FetchUsersResponse struct {
	Data UserData `json:"data,omitempty"`
}

// UserData represents the GraphQL's user data structure
type UserData struct {
	Search SearchUser `json:"search,omitempty"`
}

// SearchUser represents the GraphQL's search user data structure
type SearchUser struct {
	UserCount int64          `json:"userCount,omitempty"`
	PageInfo  model.PageInfo `json:"pageInfo,omitempty"`
	Edges     []UserEdge     `json:"edges,omitempty"`
}

// UserEdge represents the GraphQL's user edge data structure
type UserEdge struct {
	Cursor string     `json:"cursor,omitempty"`
	Node   model.User `json:"node,omitempty"`
}

// FetchReposResponse represents the GraphQL's repo data structure
type FetchReposResponse struct {
	Data RepoData `json:"data,omitempty"`
}

// RepoData represents the GraphQL's repo data structure
type RepoData struct {
	Search SearchRepo `json:"search,omitempty"`
}

// SearchRepo represents the GraphQL's search repo data structure
type SearchRepo struct {
	RepositoryCount int64          `json:"repositoryCount,omitempty"`
	PageInfo        model.PageInfo `json:"pageInfo,omitempty"`
	Edges           []RepoEdge     `json:"edges,omitempty"`
}

// RepoEdge represents the GraphQL's repo edge data structure
type RepoEdge struct {
	Cursor string     `json:"cursor,omitempty"`
	Node   model.Repo `json:"node,omitempty"`
}
