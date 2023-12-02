package cache

import (
	"sync"
	"wb/models"
)

type Cache struct {
	cache map[string]models.Order
	mutex sync.Mutex
}

func New() Cache {
	return Cache{cache: make(map[string]models.Order)}
}

func (c *Cache) Save(order models.Order) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[order.OrderUID] = order

	return nil
}

func (c *Cache) Load(uid string) (models.Order, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	order := c.cache[uid]
	if len(order.OrderUID) == 0 {
		return order, false
	}
	return order, true
}
