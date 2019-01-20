// Heavily inspired by https://github.com/pete911/examples-redigo
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

// CacheClient - interface for cache client
type CacheClient interface {
	Exists(key string) (bool, error)
	Get(key string) ([]byte, error)
	GetParsed(key string, result interface{}) error
	Set(key string, value []byte) error
}

// RedisClient - Redis cache client and wrappers
type RedisClient struct {
	pool *redis.Pool
}

// NewRedisClient - return new RedisClient
func NewRedisClient() RedisClient {
	p := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", os.Getenv("REDIS_URL"))
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

	fmt.Println("Redis cache connected")
	return RedisClient{pool: p}
}

// Ping - ping server
func (c RedisClient) Ping() error {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

// Exists - true if key exists
func (c RedisClient) Exists(key string) (bool, error) {
	conn := c.pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
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
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

// GetParsed - get item from Redis and parse it
func (c RedisClient) GetParsed(key string, result interface{}) error {
	data, err := c.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &result)
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
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}
