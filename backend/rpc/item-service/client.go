package item

import (
	"context"

	pb "github.com/lucamatsumoto/go-sales-manager/backend/rpc/item-service/proto/item"
	"github.com/micro/go-micro/client"
)

type Item struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Category string `json:"Category"`
}

type Client struct {
	itemService pb.ItemServiceClient
}

func NewClient() (*Client, error) {
	c := pb.NewItemServiceClient("go.micro.srv.item", client.DefaultClient)
	return &Client{c}, nil
}

func (c *Client) Create(ctx context.Context, name string) (*Item, error) { //also return JSON struct of response needed -> check the necessary requests
	r, err := c.itemService.CreateItem(ctx, &pb.Item{}) //fill in user struct later -> transform JSON
	if err != nil {
		return nil, err
	}
	return &Item{
		ID:       r.Item.Id,
		Name:     r.Item.Name,
		Owner:    r.Item.Owner,
		Category: r.Item.Category,
	}, nil
} //might not need to return anything

func (c *Client) Get(ctx context.Context, id string) (*Item, error) {
	r, err := c.itemService.GetItemsByName(ctx, &pb.GetItemByNameRequest{Id: id}) //put something in here
	if err != nil {
		return nil, err
	}
	return &Item{
		ID:       r.Item.Id,
		Name:     r.Item.Name,
		Owner:    r.Item.Owner,
		Category: r.Item.Category,
	}, nil
}

func (c *Client) GetAll(ctx context.Context) ([]Item, error) {
	r, err := c.itemService.GetItems(ctx, &pb.Request{})
	if err != nil {
		return nil, err
	}
	items := []Item{}
	for _, i := range r.Items {
		items = append(items, Item{
			ID:       i.Id,
			Name:     i.Name,
			Owner:    i.Owner,
			Category: i.Category,
		})
	}
	return items, nil
}

//json Wrapper implementation necessary to massage responses from gRPC over to JSON when UI calls it
