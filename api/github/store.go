package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alextanhongpin/go-github-scraper/internal/util"
)

// Store represents the interface for the Github Service
type Store interface {
	FetchUsers(location, start, end, cursor string) (*FetchUsersResponse, error)
	FetchRepos(login, start, end, cursor string) (*FetchReposResponse, error)
}

// store holds the store configuration
type store struct {
	client   *http.Client
	token    string
	endpoint string
}

// New returns a new store
func New(token, endpoint string) Store {
	return &store{
		client:   util.NewHTTPClient(),
		token:    token,
		endpoint: endpoint,
	}
}

func (s *store) FetchUsers(location, start, end, cursor string) (*FetchUsersResponse, error) {
	body := FetchUsersRequest(location, start, end, cursor, 10)
	jsonBytes, err := json.Marshal(GraphQLQuery{body.String()})
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

func (s *store) FetchRepos(login, start, end, cursor string) (*FetchReposResponse, error) {
	body := FetchReposRequest(login, start, end, cursor, 10)
	jsonBytes, err := json.Marshal(GraphQLQuery{body.String()})
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
