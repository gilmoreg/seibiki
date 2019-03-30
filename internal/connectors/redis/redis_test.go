package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestRedisDriver(t *testing.T) {
	client := New("redis://localhost:6379", newTestLogger())

	t.Run("PING", func(t *testing.T) {
		err := client.(*redisClient).Ping()
		assert.Nil(t, err)
	})

	t.Run("SET", func(t *testing.T) {
		err := client.Set("test", []byte("test"))
		assert.Nil(t, err)
	})

	t.Run("GET", func(t *testing.T) {
		res, err := client.Get("test")
		assert.Nil(t, err)
		assert.Equal(t, "test", string(res))
	})
}

func TestRedisDriverErrors(t *testing.T) {
	client := New("redis://nowhere", newTestLogger())

	t.Run("PING error", func(t *testing.T) {
		err := client.(*redisClient).Ping()
		assert.NotNil(t, err)
	})

	t.Run("SET error", func(t *testing.T) {
		err := client.Set("test", []byte("test"))
		assert.NotNil(t, err)
	})

	t.Run("GET error", func(t *testing.T) {
		_, err := client.Get("test")
		assert.NotNil(t, err)
	})
}

func newTestLogger() *zap.SugaredLogger {
	return zap.NewExample().Sugar()
}
