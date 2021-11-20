package cache

import "github.com/bradfitz/gomemcache/memcache"

type Cache struct {
	Client *memcache.Client
}

func Get() (*Cache, error) {
	cache := memcache.New("localhost:11211")
	if err := cache.Ping(); err != nil {
		return nil, err
	}

	return &Cache{Client: cache}, nil
}

func (c *Cache) Close() {
	// ugly hack but no reasonable solution was found :(
	// client will be destroyed by GC
	c.Client = nil
}
