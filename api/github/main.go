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
	FetchUsers(start, end, cursor string) (*FetchUserResponse, error)
}

// Service holds the service configuration
type service struct {
	client   *http.Client
	token    string
	country  string
	endpoint string
}

// New returns a new service
func New(token, endpoint, country string) Service {
	return &service{
		client:   util.NewHTTPClient(),
		token:    token,
		endpoint: endpoint,
		country:  country,
	}
}

func (s *service) FetchUsers(start, end, cursor string) (*FetchUserResponse, error) {
	jsonReq := makeUserPayload(s.country, start, end, cursor)
	jsonResp, err := graphqlService(s.client, s.token, s.endpoint, jsonReq)
	if err != nil {
		return nil, err
	}
	var resp FetchUserResponse
	if err := json.Unmarshal(jsonResp, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

func makeUserPayload(location, start, end, cursor string) []byte {
	q := fmt.Sprintf(`
		query {
			search(query: "location:%s created:%s..%s", type: USER, last: 10, after: %s) {
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
		location, start, end, cursor)

	out, err := json.Marshal(model.GraphQLQuery{q})
	if err != nil {
		log.Println(err)
	}
	return out
}

type FetchUserResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Search UserSearch `json:"search"`
}

type UserSearch struct {
	UserCount int64          `json:"userCount"`
	PageInfo  model.PageInfo `json:"pageInfo"`
	Edges     []UserEdge     `json:"edges"`
}

type UserEdge struct {
	Cursor string     `json:"cursor"`
	Node   model.User `json:"node"`
}
