package mongodb_test

import (
	"encoding/json"
	"testing"

	"github.com/gilmoreg/seibiki/internal/connectors/mongodb"
	"github.com/gilmoreg/seibiki/internal/dictionary"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestMongoDBDriver(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		client, err := mongodb.New("mongodb://reader:password@localhost:27017/jedict", newTestLogger())
		assert.Nil(t, err)
		res, err := client.Get(makePipeline("寒い"))
		assert.Nil(t, err)
		var entries []dictionary.Entry
		err = json.Unmarshal(res, &entries)
		// Assert we are able to deserialize this to the correct type of response
		assert.Nil(t, err)
		assert.NotNil(t, entries)
	})

	t.Run("GET error", func(t *testing.T) {
		client, err := mongodb.New("mongodb://reader:password@localhost:27017/jedict", newTestLogger())
		assert.Nil(t, err)
		_, err = client.Get(make(chan int))
		assert.NotNil(t, err)
	})

	t.Run("Connection Error", func(t *testing.T) {
		_, err := mongodb.New("", newTestLogger())
		assert.NotNil(t, err)
	})
}

func newTestLogger() *zap.SugaredLogger {
	return zap.NewExample().Sugar()
}

func makePipeline(query string) interface{} {
	return bson.M{
		"$or": bson.A{
			bson.M{"readings": query},
			bson.M{"kanji": query},
		},
	}
}
