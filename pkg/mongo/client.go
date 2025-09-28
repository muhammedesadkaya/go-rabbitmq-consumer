package mongo

import (
	"github.com/globalsign/mgo"
	"time"
)

type Client struct {
	Session *mgo.Session
}

type MongoDBClient interface {
	NewSession() *mgo.Session
	EnsureIndex(fields []string, isUnique bool, indexName string, database string, collection string) error
}

func NewClient(connectionString string, timeout time.Duration) (MongoDBClient, error) {
	session, err := mgo.DialWithTimeout(connectionString, timeout)

	if err != nil {
		return nil, err
	}

	return &Client{Session: session}, nil
}

func (c *Client) NewSession() *mgo.Session {
	newSession := c.Session.Copy()
	newSession.SetMode(mgo.Strong, true)
	return newSession
}

func (c *Client) EnsureIndex(fields []string, isUnique bool, indexName, database, collection string) error {
	localSession := c.NewSession()
	defer localSession.Close()

	index := mgo.Index{
		Key:        fields,
		Unique:     isUnique,
		Name:       indexName,
		Background: true,
	}

	col := localSession.DB(database).C(collection)

	if err := col.EnsureIndex(index); err != nil {
		return err
	}

	return nil
}
