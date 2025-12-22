package cache

import "fmt"

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
}
type SimpleCache struct {
	data map[string]string
}

func (c *SimpleCache) Get(key string) (string, error) {
	value, exists := c.data[key]
	if !exists {
		return "", fmt.Errorf("key not found")
	}
	return value, nil
}
func (c *SimpleCache) Set(key, value string) error {
	c.data[key] = value
	return nil
}

func (c *SimpleCache) Delete(key string) error {
	delete(c.data, key)
	return nil
}

func NewSimpleCache() Cache {
	return &SimpleCache{
		data: make(map[string]string),
	}
}
