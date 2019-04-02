package dictionary

import (
	"encoding/json"
	"log"

	"github.com/gilmoreg/seibiki/internal/connectors/mongodb"
	"github.com/gilmoreg/seibiki/internal/connectors/redis"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/mongodb/mongo-go-driver/bson"
	"go.uber.org/zap"
)

// Repository - repository for dictionary
type Repository interface {
	Lookup(query string) ([]Entry, error)
}

type dictionary struct {
	db     mongodb.Client
	cache  redis.Client
	logger *zap.SugaredLogger
}

// New - new Dictionary Repository
func New(db mongodb.Client, cache redis.Client, logger *zap.SugaredLogger) Repository {
	return &dictionary{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

// Lookup - find entries from cache or db
func (d *dictionary) Lookup(query string) ([]Entry, error) {
	ok, cached, err := d.cacheLookup(query)
	if ok {
		return cached, nil
	}
	if err != nil {
		d.logger.Error(err)
		return nil, err
	}
	pipeline := bson.M{
		"$or": bson.A{
			bson.M{"readings": query},
			bson.M{"kanji": query},
		},
	}
	rawEntries, err := d.db.Get(pipeline)
	if err != nil {
		d.logger.Error(err)
		return nil, err
	}
	entries, err := decode(rawEntries)
	if err != nil {
		d.logger.Error(err)
		return nil, err
	}
	go d.cacheFill(query, entries)
	return entries, nil
}

func (d *dictionary) cacheLookup(query string) (bool, []Entry, error) {
	// d.logger.Infof("Checking cache for %s", query)
	data, err := d.cache.Get(query)
	if err == redigo.ErrNil {
		// d.logger.Infof("%s does not exist in cache. Fetching from db", query)
		return false, nil, nil
	}
	if err != nil {
		d.logger.Errorf("Error connecting to cache: %s", err.Error())
		return false, nil, err
	}
	var entries []Entry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		d.logger.Error(err)
		return false, nil, err
	}
	// d.logger.Infof("fetched %s", query)
	return true, entries, nil
}

func (d *dictionary) cacheFill(query string, entries []Entry) {
	bytes, err := json.Marshal(&entries)
	if err != nil {
		log.Fatal(err)
	}
	// d.logger.Infof("setting %s in cache", query)
	err = d.cache.Set(query, bytes)
	if err != nil {
		d.logger.Errorf("error setting cache: %s", err.Error())
	}
}

// decode - convert bson.A to []Entry
func decode(rawEntries []byte) ([]Entry, error) {
	var entries []Entry
	err := json.Unmarshal(rawEntries, &entries)
	return entries, err
}
