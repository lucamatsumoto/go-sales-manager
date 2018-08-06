package graph

import (
	"context"
	"errors"
)

var (
	InvalidParamsErr = errors.New("Invalid Parameter")
)

func (s *GraphQLServer) Mutation_createUser(ctx context.Context) {
	//generate account input
}
