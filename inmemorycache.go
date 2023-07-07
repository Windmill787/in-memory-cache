package inmemorycache

type Cache struct {
	storage map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(key string, value interface{}) {
	c.storage[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.storage[key]
}

func (c *Cache) Delete(key string) {
	delete(c.storage, key)
}
