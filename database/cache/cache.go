package cache

import (
	"sync"
	"wb/config"
)

type Cache struct {
	cache map[string]config.Order
	mutex sync.Mutex
}

func New() Cache {
	return Cache{cache: make(map[string]config.Order)}
}

func (c *Cache) Save(order config.Order) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[order.OrderUID] = order

	return nil
}

func (c *Cache) Load(uid string) (config.Order, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	order := c.cache[uid]
	if len(order.OrderUID) == 0 {
		return order, false
	}
	return order, true
}
