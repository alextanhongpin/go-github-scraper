package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	fetchGraphqlUser()
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        myHandler,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// log.Fatal(s.ListenAndServe())
}

func fetchGraphqlUser() error {
	url := "https://api.github.com/graphql"

	var jsonStr = []byte(`
		{
			"query": "query {
				search(query: "location:malaysia created:2018-01-01..2018-01-10", type: USER, last: 10, after: "Y3Vyc29yOjIw") {
					userCount,
					pageInfo {
						hasNextPage,
						startCursor,
						endCursor,
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
			}
		}"
	`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("Authorization", "bearer token")
	if err != nil {
		return err
	}
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	return nil
}
