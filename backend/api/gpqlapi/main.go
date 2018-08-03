package main

import (
	pb "github.com/lucamatsumoto/go-sales-manager/backend/rpc/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
)

func main() {
	pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	//schema := graphql.
	//gql.NewRouter()

	/*srv := &http.Server{
		Addr: ":8080",
	}*/ // the graphql server that we are hosting on
}
