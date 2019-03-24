package mongodb

import (
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestMongoDBDriver(t *testing.T) {
	client := New("mongodb://localhost:27017/jedict", newTestLogger())
	res, err := client.Get(makePipeline("寒い"))
	assert.Nil(t, err)
	assert.NotNil(t, res)
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
