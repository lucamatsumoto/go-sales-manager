package gql

import (
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var page = []byte("HELLO")

//the request from the client goes here

func NewRouter(schema *graphql.Schema) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	mux.Handle("/user", &relay.Handler{Schema: schema}) //we can create a routes.go class later to handle this
	mux.Handle("/item", &relay.Handler{Schema: schema})
	mux.Handle("/email", &relay.Handler{Schema: schema})
	//transfers requests to graphql over to gRPC microservices

	return mux
}
