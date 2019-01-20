package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
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
	cache      CacheClient
	client     *mongo.Client
	collection *mongo.Collection
}

// New - creat new MongoDBRepository
func (m MongoDBRepository) New(connectionString string, cc CacheClient) MongoDBRepository {
	client, err := mongo.Connect(context.TODO(), connectionString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected. Testing connection...")

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Test successful. Database connected.")

	m.client = client
	m.collection = client.Database("jedict").Collection("entries")
	m.cache = cc
	return m
}

func (m MongoDBRepository) cacheLookup(query string) (bool, []Entry, error) {
	ok, err := m.cache.Exists(query)
	if err != nil {
		return false, nil, err
	}
	if !ok {
		return false, nil, nil
	}
	var entries []Entry
	err = m.cache.GetParsed(query, entries)
	return true, entries, err
}

func (m MongoDBRepository) cacheFill(query string, entries []Entry) {
	bytes, err := json.Marshal(&entries)
	if err != nil {
		log.Fatal(err)
	}
	err = m.cache.Set(query, bytes)
	if err != nil {
		log.Fatal(err)
	}
}

// Lookup - perform a dictionary lookup
func (m MongoDBRepository) Lookup(query string) []Entry {
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
	options := options.Find()
	var results []Entry
	cur, err := m.collection.Find(context.TODO(), pipeline, options)
	defer cur.Close(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Entry
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	defer m.cacheFill(query, results)
	return results
}
