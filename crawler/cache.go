package crawler

import "sync"

type URLCache struct {
	m    *sync.Mutex
	data map[string]bool
}

func NewURLCache() *URLCache {
	return &URLCache{
		m:    &sync.Mutex{},
		data: map[string]bool{},
	}
}

func (c *URLCache) Add(url string) {
	c.m.Lock()
	defer c.m.Unlock()
	c.data[url] = true
}

func (c *URLCache) Has(url string) bool {
	c.m.Lock()
	defer c.m.Unlock()
	if _, ok := c.data[url]; !ok {
		return false
	}
	return true
}

func (c *URLCache) FetchAll() map[string]bool {
	c.m.Lock()
	defer c.m.Unlock()
	return c.data
}
