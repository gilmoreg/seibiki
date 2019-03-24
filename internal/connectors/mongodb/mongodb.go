package mongodb

import (
	"context"
	"encoding/json"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"go.uber.org/zap"
)

// Client -
type Client interface {
	Get(query interface{}) ([]byte, error)
}

type client struct {
	client *mongo.Client
	logger *zap.SugaredLogger
}

// New - create new mongodb client
func New(connectionString string, logger *zap.SugaredLogger) Client {
	c := client{logger: logger}
	c.connect(connectionString)
	return &c
}

// Connect - connect to database
func (m *client) connect(connectionString string) error {
	client, err := mongo.Connect(context.TODO(), connectionString)
	if err != nil {
		m.logger.Error(err)
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		m.logger.Error(err)
		return err
	}

	m.client = client
	return nil
}

// Lookup - perform a dictionary lookup
func (m *client) Get(query interface{}) ([]byte, error) {
	cur, err := m.client.Database("jedict").Collection("entries").Find(context.TODO(), query, options.Find())
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}
	defer cur.Close(context.TODO())

	var result bson.A

	for cur.Next(context.TODO()) {
		var elem bson.M
		err := cur.Decode(&elem)
		if err != nil {
			m.logger.Warn(err)
			continue
		}

		result = append(result, elem)
	}
	return json.Marshal(result)
}
