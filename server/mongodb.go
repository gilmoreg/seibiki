package main

import (
	"context"
	"log"
	"unicode"

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
	Lookup(query string) Entry
}

// MongoDBRepository - DictionaryRepository for MongoDB
type MongoDBRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// New - creat new MongoDBRepository
func (m MongoDBRepository) New(connectionString string) MongoDBRepository {
	client, err := mongo.Connect(context.TODO(), connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	m.client = client
	m.collection = client.Database("jedict").Collection("entries")

	return m
}

// Lookup - perform a dictionary lookup
func (m MongoDBRepository) Lookup(query string) []*Entry {
	pipeline := bson.M{
		"$or": bson.A{
			bson.M{"readings": query},
			bson.M{"kanji": query},
		},
	}
	options := options.Find()
	var results []*Entry
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

		results = append(results, &elem)
	}
	return results
}

func kanjiOnly(s string) bool {
	for _, r := range s {
		if !unicode.In(r, unicode.Ideographic) {
			return false
		}
	}
	return s != ""
}

// might be unicode.Unified_Ideograph
func kanaOnly(s string) bool {
	for _, r := range s {
		if !unicode.In(r, unicode.Ideographic) {
			return false
		}
	}
	return s != ""
}
