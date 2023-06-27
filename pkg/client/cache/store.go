package cache

import "sync"

type cache struct {
	lock  sync.RWMutex
	items map[string]interface{}
}

func (c *cache) Add(ID string, obj interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[ID] = obj
}

func (c *cache) Update(ID string, obj interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[ID] = obj
}
func (c *cache) Delete(ID string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.items, ID)
}
func (c *cache) Get(ID string) (item interface{}, exists bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	item, exists = c.items[ID]
	return
}
func (c *cache) List() []interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	result := make([]interface{}, 0, len(c.items))
	for _, item := range c.items {
		result = append(result, item)
	}
	return result
}

//func NewStore() *cache
