package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/util"
)

// Store represents the interface for the Github Service
type (
	Store interface {
		FetchUsers(req FetchUsersRequest) (*FetchUsersResponse, error)
		FetchRepos(req FetchReposRequest) (*FetchReposResponse, error)
	}

	// store holds the store configuration
	store struct {
		client   *http.Client
		token    string
		endpoint string
	}
)

// NewStore returns a new store
func NewStore(token, endpoint string) Store {
	return &store{
		client:   util.NewHTTPClient(),
		token:    token,
		endpoint: endpoint,
	}
}

func (s *store) FetchUsers(req FetchUsersRequest) (*FetchUsersResponse, error) {
	jsonBytes, err := json.Marshal(GraphQLQuery{req.String()})
	if err != nil {
		return nil, err
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

func (s *store) FetchRepos(req FetchReposRequest) (*FetchReposResponse, error) {
	jsonBytes, err := json.Marshal(GraphQLQuery{req.String()})
	if err != nil {
		return nil, err
	}

	jsonResp, err := graphqlService(s.client, s.token, s.endpoint, jsonBytes)
	if err != nil {
		return nil, err
	}

	var resp FetchReposResponse
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
