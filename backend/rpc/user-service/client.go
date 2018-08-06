package user

import (
	"context"

	pb "github.com/lucamatsumoto/go-sales-manager/backend/rpc/user-service/proto/user"
	"github.com/micro/go-micro/client"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Client struct {
	userService pb.UserServiceClient
}

func NewClient() (*Client, error) {
	c := pb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
	return &Client{c}, nil
}

func (c *Client) Create(ctx context.Context, name string) (*User, error) { //also return JSON struct of response needed -> check the necessary requests
	r, err := c.userService.Create(ctx, &pb.User{}) //fill in user struct later -> transform JSON
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    r.User.Id,
		Name:  r.User.Name,
		Email: r.User.Email,
	}, nil
} //might not need to return anything

func (c *Client) Get(ctx context.Context) (*User, error) {
	r, err := c.userService.Get(ctx, &pb.User{}) //put something in here
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    r.User.Id,
		Name:  r.User.Name,
		Email: r.User.Email,
	}, nil
}

func (c *Client) GetAll(ctx context.Context) ([]User, error) {
	r, err := c.userService.GetAll(ctx, &pb.Request{})
	if err != nil {
		return nil, err
	}
	users := []User{}
	for _, usr := range r.Users {
		users = append(users, User{
			ID:    usr.Id,
			Name:  usr.Name,
			Email: usr.Email,
		})
	}
	return users, nil
}

//json Wrapper implementation necessary to massage responses from gRPC over to JSON when UI calls it
