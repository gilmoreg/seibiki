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
	client := mongodb.New("mongodb://reader:password@localhost:27017/jedict", newTestLogger())
	res, err := client.Get(makePipeline("寒い"))
	assert.Nil(t, err)
	var entries []dictionary.Entry
	err = json.Unmarshal(res, &entries)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(entries))
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
