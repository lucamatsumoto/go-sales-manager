package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/lucamatsumoto/go-sales-manager/backend/rpc/user-service/proto/user"

	"github.com/micro/go-micro"
)

const (
	host = "localhost:27017"
)

func main() {
	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		dbHost = host
	}

	session, err := CreateSession(host)

	defer session.Close()
	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to database with host %s - %v", host, err)
	}

	repo := &UserRepository{session}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	publisher := micro.NewPublisher("user.created", srv.Client())

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, publisher})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
