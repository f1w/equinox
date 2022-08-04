package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v9"
)

type Cache struct {
	store Store
	TTL   time.Duration
}

type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl time.Duration) error
	Delete(key string) error
	Clear() error
}

var (
	ErrCacheIsDisabled = errors.New("Cache is disabled")
)

// Creates a new Cache using BigCache.
//
// Requires a BigCache config that can be created with bigcache.DefaultConfig(n*time.Minute).
func NewBigCache(config bigcache.Config) (*Cache, error) {
	bigcache, err := bigcache.NewBigCache(config)

	if err != nil {
		return nil, err
	}

	cache := &Cache{
		store: &BigCacheStore{
			client: bigcache,
		},
		TTL: config.LifeWindow,
	}

	return cache, nil
}

// Creates a new Cache using go-redis.
func NewRedis(ctx context.Context, options *redis.Options, ttl time.Duration) (*Cache, error) {
	if options == nil {
		return nil, fmt.Errorf("redis options is empty")
	}

	redis := redis.NewClient(options)

	err := redis.Ping(ctx).Err()

	if err != nil {
		return nil, err
	}

	cache := &Cache{
		store: &RedisStore{
			client: redis,
			ttl:    ttl,
			ctx:    ctx,
		},
		TTL: ttl,
	}

	return cache, nil
}

// Returns an item from the cache. If no item is found, returns nil for the item and error.
func (c *Cache) Get(key string) ([]byte, error) {
	if c.TTL == 0 {
		return nil, ErrCacheIsDisabled
	}

	value, err := c.store.Get(key)

	if err != nil {
		return nil, err
	}

	return value, nil
}

// Saves an item under the key provided.
func (c *Cache) Set(key string, item []byte) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}

	return c.store.Set(key, item, c.TTL)
}

// Deletes an item from the cache.
func (c *Cache) Delete(key string) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}

	return c.store.Delete(key)
}

// Clears the entire cache.
func (c *Cache) Clear() error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}

	return c.store.Clear()
}
