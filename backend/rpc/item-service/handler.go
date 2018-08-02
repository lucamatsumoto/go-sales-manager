package main

import (
	"golang.org/x/net/context"

	pb "github.com/lucamatsumoto/go-sales-manager/backend/rpc/item-service/proto/item"
	mgo "gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &ItemRepository{s.session.Clone()}
}

func (s *service) CreateItem(ctx context.Context, req *pb.Item, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	//add ID logic here later when we create our client

	err := repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Item = req
	return nil
}

func (s *service) GetItems(ctx context.Context, req *pb.Request, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	items, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Items = items
	return nil
}

func (s *service) GetItemsByName(ctx context.Context, req string, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	items, err := repo.GetItemsByName(req)
	if err != nil {
		return err
	}
	res.Items = items
	return nil
}
