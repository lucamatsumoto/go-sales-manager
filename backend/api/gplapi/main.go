package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/lucamatsumoto/go-sales-manager/backend/api/gpqlapi/server"
)

func main() {
	serv, err := server.NewGraphQLServer()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", handler.GraphQL())
}
