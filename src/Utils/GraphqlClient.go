package Utils

import (
	"context"

	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

func CreateClient(token string) *graphql.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	Client := graphql.NewClient("https://api-dev.panorra.com/v1/graphql", httpClient)
	return Client
}
