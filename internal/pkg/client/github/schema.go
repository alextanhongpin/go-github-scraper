package github

import (
	"fmt"
)

// GraphQLQuery represents the structure for the Github's GraphQL API calls
type GraphQLQuery struct {
	Query string `json:"query"`
}

// PageInfo represents the pagination structure from Github
type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

// FetchUsersRequest represents the request for the GraphQL service
type FetchUsersRequest struct {
	Location string
	Start    string
	End      string
	Cursor   string
	Limit    int
}

func (f FetchUsersRequest) String() string {
	var cursor string
	if f.Cursor == "" {
		cursor = ""
	} else {
		// String must be quoted for the graphql api to work
		cursor = fmt.Sprintf(`"%s"`, f.Cursor)
	}
	return fmt.Sprintf(`
		query {
			search(query: "location:%s created:%s..%s", type: USER, first: %d, after: %s) {
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
							websiteUrl,
							repositories(last: 0) {
								totalCount
							},
							gists(last: 0) {
								totalCount
							},
							followers(last: 0) {
								totalCount
							},
							following(last: 0) {
								totalCount
							},
						}
					}
				}
			}
		}`, f.Location, f.Start, f.End, f.Limit, cursor)
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
	UserCount int64      `json:"userCount,omitempty"`
	PageInfo  PageInfo   `json:"pageInfo,omitempty"`
	Edges     []UserEdge `json:"edges,omitempty"`
}

// UserEdge represents the GraphQL's user edge data structure
type UserEdge struct {
	Cursor string `json:"cursor,omitempty"`
	Node   User   `json:"node,omitempty"`
}

// FetchReposRequest represents the request for the GraphQL service
type FetchReposRequest struct {
	Login  string
	Start  string
	End    string
	Cursor string
	Limit  int
}

func (f FetchReposRequest) String() string {
	var cursor string
	if f.Cursor == "" {
		cursor = ""
	} else {
		cursor = fmt.Sprintf(`"%s"`, f.Cursor)
	}
	return fmt.Sprintf(`query {
		search(query: "user:%s created:%s..%s", type: REPOSITORY, first: %d, after: %s) {
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
						},
						url
					}
				}
			}
		}
	}`, f.Login, f.Start, f.End, f.Limit, cursor)
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
	RepositoryCount int64      `json:"repositoryCount,omitempty"`
	PageInfo        PageInfo   `json:"pageInfo,omitempty"`
	Edges           []RepoEdge `json:"edges,omitempty"`
}

// RepoEdge represents the GraphQL's repo edge data structure
type RepoEdge struct {
	Cursor string `json:"cursor,omitempty"`
	Node   Repo   `json:"node,omitempty"`
}
