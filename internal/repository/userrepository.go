package repository

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/go-rabbitmq-consumer-app/internal/repository/entity"
	"gitlab.com/go-rabbitmq-consumer-app/pkg/mongo"
)

const (
	userCollectionName = "users"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetByFullName(ctx context.Context, fullName string) (*entity.User, error)
	Insert(ctx context.Context, user *entity.User) error
	Upsert(ctx context.Context, user *entity.User) error
}

type userRepository struct {
	mongoClient  mongo.MongoDBClient
	databaseName string
}

func (self *userRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	var session = self.mongoClient.NewSession()
	defer session.Close()

	var records []*entity.User

	err := session.
		DB(self.databaseName).
		C(userCollectionName).
		Find(nil).All(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (self *userRepository) GetByFullName(ctx context.Context, fullName string) (*entity.User, error) {
	var session = self.mongoClient.NewSession()
	defer session.Close()

	var record *entity.User

	err := session.
		DB(self.databaseName).
		C(userCollectionName).
		Find(bson.M{
			"FullName": fullName,
		}).One(&record)

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (self *userRepository) Insert(ctx context.Context, user *entity.User) error {
	var session = self.mongoClient.NewSession()
	defer session.Close()

	err := session.
		DB(self.databaseName).
		C(userCollectionName).
		Insert(user)

	if err != nil {
		return err
	}

	return nil
}

func (self *userRepository) Upsert(ctx context.Context, user *entity.User) error {
	var session = self.mongoClient.NewSession()
	defer session.Close()

	_, err := session.
		DB(self.databaseName).
		C(userCollectionName).
		Upsert(bson.M{"FullName": user.FullName}, user)

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(client mongo.MongoDBClient, databaseName string) UserRepository {

	if err := client.EnsureIndex([]string{"CreateDate"}, true, "CreateDate", databaseName, userCollectionName); err != nil {
		fmt.Println(err)
	}

	return &userRepository{mongoClient: client, databaseName: databaseName}
}
