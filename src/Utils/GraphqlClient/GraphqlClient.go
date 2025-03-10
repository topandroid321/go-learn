package GraphqlClient

import (
	"context"
	"net/http"

	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

// HEADER USER
func CreateClient(token string) *graphql.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	Client := graphql.NewClient("https://api-dev.panorra.com/v1/graphql", httpClient)
	return Client
}

// HEADER ADMIN
func CreateAdmin() *graphql.Client {
	headers := http.Header{}
	headers.Set("x-hasura-admin-secret", "ZxWN8eNMC7Wnugtb") // Replace with your actual admin secret

	// Create an HTTP client with the custom headers
	client := &http.Client{
		Transport: &headerTransport{Header: headers},
	}

	// Create a GraphQL client with a custom HTTP client
	return graphql.NewClient("https://api-dev.panorra.com/v1/graphql", client)
}

// headerTransport is a custom transport that adds headers to HTTP requests.
type headerTransport struct {
	Header http.Header
}

// RoundTrip adds the specified headers to the request before sending it.
func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add headers to the request
	for key, values := range t.Header {
		req.Header[key] = values
	}

	// Perform the request
	return http.DefaultTransport.RoundTrip(req)
}
