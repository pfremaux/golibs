package cache

import "fmt"

type SimpleCache[V any] struct {
	data map[string]V
}

func (c *SimpleCache[V]) Get(key string) (V, error) {
	value, exists := c.data[key]
	if !exists {
		var zero V
		return zero, fmt.Errorf("key not found")
	}
	return value, nil
}
func (c *SimpleCache[V]) Set(key string, value V) error {
	c.data[key] = value
	return nil
}

func (c *SimpleCache[V]) Delete(key string) error {
	delete(c.data, key)
	return nil
}

func (c *SimpleCache[V]) Values() ([]V, error) {
	values := make([]V, 0, len(c.data))
	for _, v := range c.data {
		values = append(values, v)
	}
	return values, nil
}

func NewSimpleCache[V any]() Cache[string, V] {
	return &SimpleCache[V]{
		data: make(map[string]V),
	}
}
