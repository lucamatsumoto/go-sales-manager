package main

import (
	"errors"
	"fmt"

	pb "github.com/lucamatsumoto/go-sales-manager/backend/rpc/user-service/proto/user"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName         = "fill"
	userCollection = "users"
)

type Repository interface {
	Create(*pb.User) error
	Get(string) (*pb.User, error)
	GetByEmail(string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	Close()
}

type UserRepository struct {
	session *mgo.Session
}

func (repo *UserRepository) Create(user *pb.User) error {
	err := ensureIndex(repo.collection())
	if err != nil {
		return err
	}
	err = repo.collection().Insert(user)
	if err != nil {
		if mgo.IsDup(err) {
			return errors.New(fmt.Sprintf("Email %s already exists", user.Email))
		} //check for duplicate emails
	}
	return err
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var userFound *pb.User
	err := repo.collection().Find(nil).Select(bson.M{"id": id}).One(userFound)
	return userFound, err
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	var user *pb.User
	err := repo.collection().Find(nil).Select(bson.M{"email": email}).One(&user)
	return user, err
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User

	err := repo.collection().Find(nil).All(&users)
	return users, err
}

func (repo *UserRepository) Close() {
	repo.session.Close()
}

func (repo *UserRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(userCollection)
}

func ensureIndex(collection *mgo.Collection) error {
	index := mgo.Index{
		Key:      []string{"email"},
		Unique:   true,
		DropDups: true,
	}
	return collection.EnsureIndex(index)
}
