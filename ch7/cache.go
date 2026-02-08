package ch7
type Cache(K Comparable, V any) struct {
	data map[K]v
}

func NewCache[K Comparable, V any] () Cache[K, V] {
	return Cache[K,V] {
		data : make(map[K]V),
	}
}

func (c *Cache[K,V]) Read(key K) (V, bool) {
	v, found := c.data[key]
	return v, found
}

func(c *Cache[K,V]) Upsert(key K, value V) error {
	c.data[key] = value
	return nil
}

func (c *Cache[K,V]) Delete(key K) error {
	delete(c.data, key)
}