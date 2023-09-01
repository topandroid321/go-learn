package Utils

import "github.com/hasura/go-graphql-client"

var Client = graphql.NewClient("https://example.com/graphql", nil)
