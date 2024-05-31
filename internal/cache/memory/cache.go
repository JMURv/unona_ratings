package memory

import (
	"context"
	errs "github.com/JMURv/unona/ratings/internal/cache"
	md "github.com/JMURv/unona/ratings/pkg/model"
	"sync"
	"time"
)

type Cache struct {
	sync.Mutex
	data map[string]*md.Rating
}

func New() *Cache {
	return &Cache{data: make(map[string]*md.Rating)}
}

func (c *Cache) Get(_ context.Context, key string) (*md.Rating, error) {
	c.Lock()
	defer c.Unlock()
	if v, ok := c.data[key]; !ok {
		return nil, errs.ErrNotFoundInCache
	} else {
		return v, nil
	}
}

func (c *Cache) Set(_ context.Context, t time.Duration, key string, r *md.Rating) error {
	c.Lock()
	defer c.Unlock()
	c.data[key] = r
	return nil
}

func (c *Cache) Delete(_ context.Context, key string) error {
	c.Lock()
	defer c.Unlock()
	delete(c.data, key)
	return nil
}
