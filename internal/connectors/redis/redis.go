// Package redis -
// Heavily inspired by https://github.com/pete911/examples-redigo
package redis

import (
	"errors"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

// Client - interface for cache client
type Client interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

type redisClient struct {
	pool   *redis.Pool
	logger *zap.SugaredLogger
}

// New - return new redisClient
func New(l *zap.SugaredLogger) Client {
	p := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(os.Getenv("REDIS_URL"))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return &redisClient{pool: p, logger: l}
}

// Ping - ping server
func (c redisClient) Ping() error {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		c.logger.Error(err)
		return err
	}
	return nil
}

// Get - get from Redis
func (c redisClient) Get(key string) ([]byte, error) {
	conn := c.pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, err
	}
	return data, err
}

// Set - set in Redis
func (c redisClient) Set(key string, value []byte) error {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		c.logger.Error(v)
		return errors.New(v)
	}
	return err
}
