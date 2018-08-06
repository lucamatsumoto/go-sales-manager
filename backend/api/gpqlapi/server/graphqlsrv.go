//go:generate gqlgen -schema ../../gql/schema.graphql
package server

import (
	"github.com/lucamatsumoto/go-sales-manager/backend/rpc/item-service"
	"github.com/lucamatsumoto/go-sales-manager/backend/rpc/user-service"
)

type GraphQLServer struct {
	userClient *user.Client
	itemClient *item.Client
}

func NewGraphQLServer() (*GraphQLServer, error) {
	userClient, err := user.NewClient()
	if err != nil {
		return nil, err
	}

	itemClient, err := item.NewClient()
	if err != nil {
		return nil, err
	}
	return &GraphQLServer{
		userClient,
		itemClient,
	}, nil
}
