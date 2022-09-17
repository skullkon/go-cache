package async_cache

import (
	"async-cache/rw_mutex_cache"
	"context"
	"errors"
)

var ErrTimeout = errors.New("timeout error")
var ErrNotFound = errors.New("value not found")

type AsyncCache struct {
	c *rw_mutex_cache.MutexStorage
}

func NewAsyncCache() *AsyncCache {
	return &AsyncCache{c: rw_mutex_cache.New()}
}
func (c *AsyncCache) Get(ctx context.Context, key string) (string, error) {
	ch := make(chan string)
	go func() {
		defer close(ch)
		v, ok := c.c.Get(key)
		if ok {
			ch <- v
		}
	}()

	select {
	case <-ctx.Done():
		return "", ErrTimeout

	case x, ok := <-ch:
		if ok {
			return x, nil
		}
		return "", ErrNotFound
	}
}

func (c *AsyncCache) Add(ctx context.Context, key, value string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.c.Set(key, value)
	}()

	select {
	case <-ctx.Done():
		return ErrTimeout
	case <-ch:
		return nil
	}
}
