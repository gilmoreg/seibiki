package service

import (
	"testing"

	"github.com/gilmoreg/seibiki/internal/connectors/mongodb"
	"github.com/gilmoreg/seibiki/internal/connectors/redis"
	"github.com/gilmoreg/seibiki/internal/dictionary"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestService(t *testing.T) {
	testService := createTestService()
	res := testService.Lookup("飲む")
	assert.NotNil(t, res)
}

func createTestService() LookupService {
	log := zap.NewExample().Sugar()
	c := redis.New("redis://localhost:6379", log)
	m, err := mongodb.New("mongodb://reader:password@localhost:27017/jedict", log)
	if err != nil {
		panic(err)
	}
	d := dictionary.New(m, c, log)
	return New(log, d)
}
