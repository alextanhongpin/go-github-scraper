package model

// GraphQLQuery represents the structure for the Github's GraphQL API calls
type GraphQLQuery struct {
	Query string `json:"query"`
}
