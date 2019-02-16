package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"go.uber.org/zap"
)

// Entry - dictionary entry
type Entry struct {
	Sequence     int      `json:"sequence"`
	Kanji        []string `json:"kanji"`
	Readings     []string `json:"readings"`
	Meanings     []string `json:"meanings"`
	PartOfSpeech string   `json:"partofspeech"`
}

// DictionaryRepository - repository for dictionary
type DictionaryRepository interface {
	Lookup(query string) []Entry
}

// MongoDBRepository - DictionaryRepository for MongoDB
type MongoDBRepository struct {
	cache  CacheClient
	client *mongo.Client
	logger *zap.SugaredLogger
}

// Connect - connect to database
func (m *MongoDBRepository) Connect(connectionString string) {
	client, err := mongo.Connect(context.TODO(), connectionString)
	if err != nil {
		log.Fatal(err)
	}

	m.logger.Info("Database connected. Testing connection...")

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	m.client = client
	m.logger.Info("Test successful. Database connected.")
}

func (m *MongoDBRepository) cacheLookup(query string) (bool, []Entry, error) {
	m.logger.Infof("Checking cache for %s", query)
	ok, err := m.cache.Exists(query)
	if err != nil {
		return false, nil, err
	}
	if !ok {
		m.logger.Infof("%s does not exist in cache", query)
		return false, nil, nil
	}
	m.logger.Infof("%s exists in cache. Fetching...", query)
	data, err := m.cache.Get(query)
	if err != nil {
		return false, nil, err
	}
	var entries []Entry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return false, nil, err
	}
	m.logger.Infof("fetched %s", query)
	return true, entries, nil
}

func (m *MongoDBRepository) cacheFill(query string, entries []Entry) {
	bytes, err := json.Marshal(&entries)
	if err != nil {
		log.Fatal(err)
	}
	m.logger.Infof("setting %s in cache", query)
	err = m.cache.Set(query, bytes)
	if err != nil {
		log.Fatal(err)
	}
}

// Lookup - perform a dictionary lookup
func (m *MongoDBRepository) Lookup(query string) []Entry {
	ok, entries, err := m.cacheLookup(query)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return entries
	}

	pipeline := bson.M{
		"$or": bson.A{
			bson.M{"readings": query},
			bson.M{"kanji": query},
		},
	}
	var results []Entry
	cur, err := m.client.Database("jedict").Collection("entries").Find(context.TODO(), pipeline, options.Find())
	defer cur.Close(context.TODO())
	if err != nil {
		m.logger.Error(err)
		return make([]Entry, 0)
	}

	for cur.Next(context.TODO()) {
		var elem Entry
		err := cur.Decode(&elem)
		if err != nil {
			m.logger.Error(err)
			continue
		}

		results = append(results, elem)
	}
	defer m.cacheFill(query, results)
	return results
}
