package cache

import (
	"errors"

	"github.com/allegro/bigcache/v3"
)

type BigCacheClient interface {
	Get(key string) ([]byte, error)
	Set(key string, entry []byte) error
	Delete(key string) error
	Reset() error
}

type BigCacheStore struct {
	client BigCacheClient
}

func (s *BigCacheStore) Get(key string) ([]byte, error) {
	item, err := s.client.Get(key)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, nil
	}
	return item, err
}

func (s *BigCacheStore) Set(key string, value []byte) error {
	return s.client.Set(key, value)
}

func (s *BigCacheStore) Delete(key string) error {
	return s.client.Delete(key)
}

func (s *BigCacheStore) Clear() error {
	return s.client.Reset()
}
