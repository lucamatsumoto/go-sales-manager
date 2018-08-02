package main

import (
	"context"
	"log"
	"net/smtp"

	pb "github.com/lucamatsumoto/go-sales-manager/backend/rpc/user-service/proto/user"
	micro "github.com/micro/go-micro"
)

const topic = "user.created"

const email = "lucamatsumoto@gmail.com"

type Template struct {
	Name string
	URL  string
}

var auth smtp.Auth

type Subscriber struct{}

func sendEmail(user *pb.User) error {
	auth = smtp.PlainAuth("", email, "password", "smtp.gmail.com")
	template := Template{
		Name: user.Name,
		URL:  "someURL",
	}

	req := CreateNewRequest(
		[]string{user.Name},
		"Welcome!",
		"Welcome to our website",
	)

	err := req.ParseTemplate("emailTemplate.html", template)

	if err != nil {
		return err
	}

	ok, err := req.SendEmail(auth, email)
	if err != nil {
		log.Printf("Email Sent: %v", ok)
		return err
	}
	return nil
}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	go sendEmail(user)
	log.Printf("Sending email to %s", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))
}
