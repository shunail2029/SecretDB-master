package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection ...
type Connection struct {
	clt    *mongo.Client
	dbname string
}

// newConnection is a constructor of Connection
func newConnection(ctx context.Context) *Connection {
	c := new(Connection)

	err := c.connect(ctx)
	if err != nil {
		return nil
	}
	c.dbname = "secretdb"

	return c
}

// create connection to local database
func (c *Connection) connect(ctx context.Context) error {
	var err error
	c.clt, err = mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return err
	}
	return nil
}

// get database
func (c *Connection) db() *mongo.Database {
	return c.clt.Database(c.dbname)
}

// get collection
func (c *Connection) collection(name string) *mongo.Collection {
	return c.db().Collection(name)
}
