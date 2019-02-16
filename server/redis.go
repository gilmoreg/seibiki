// Heavily inspired by https://github.com/pete911/examples-redigo
package main

import (
	"os"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

// CacheClient - interface for cache client
type CacheClient interface {
	Exists(key string) (bool, error)
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

// RedisClient - Redis cache client and wrappers
type RedisClient struct {
	logger *zap.SugaredLogger
	pool   *redis.Pool
}

// NewRedisClient - return new RedisClient
func NewRedisClient(l *zap.SugaredLogger) RedisClient {
	p := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			redisURL := os.Getenv("REDIS_URL")
			redisURL = strings.Replace(redisURL, "redis://", "", -1)
			c, err := redis.Dial("tcp", redisURL)
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

	l.Info("Redis cache connected")
	return RedisClient{pool: p, logger: l}
}

// Ping - ping server
func (c RedisClient) Ping() error {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		c.logger.Errorf("cannot 'PING' db: %v", err)
		return err
	}
	return nil
}

// Exists - true if key exists
func (c RedisClient) Exists(key string) (bool, error) {
	conn := c.pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		c.logger.Errorf("error checking if key %s exists: %v", key, err)
		return ok, err
	}
	return ok, err
}

// Get - get from Redis
func (c RedisClient) Get(key string) ([]byte, error) {
	conn := c.pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		c.logger.Errorf("error getting key %s: %v", key, err)
		return data, err
	}
	return data, err
}

// Set - set in Redis
func (c RedisClient) Set(key string, value []byte) error {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		c.logger.Errorf("error setting key %s to %s: %v", key, v, err)
		return err
	}
	return err
}
