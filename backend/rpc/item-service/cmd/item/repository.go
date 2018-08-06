package main

import (
	pb "github.com/lucamatsumoto/go-sales-manager/backend/rpc/item-service/proto/item"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName         = "items"
	itemCollection = "items"
)

type Repository interface {
	Create(*pb.Item) error
	GetAll() ([]*pb.Item, error)
	GetItemsByName(string) ([]*pb.Item, error)
	Close()
}

type ItemRepository struct {
	session *mgo.Session
}

func (repo *ItemRepository) Create(item *pb.Item) error {
	return repo.collection().Insert(item)
}

func (repo *ItemRepository) GetAll() ([]*pb.Item, error) {
	var items []*pb.Item
	err := repo.collection().Find(nil).All(&items)
	return items, err
}

func (repo *ItemRepository) GetItemsByName(name string) ([]*pb.Item, error) {
	var items []*pb.Item
	err := repo.collection().Find(nil).Select(bson.M{"name": name}).All(&items)
	return items, err
}

func (repo *ItemRepository) Close() {
	repo.session.Close()
}

func (repo *ItemRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(itemCollection)
}
