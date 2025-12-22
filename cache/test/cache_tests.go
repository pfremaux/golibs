package test

import (
	"testing"

	"github.com/cache/pkg/cache"
)

func TestCache(t *testing.T) {
	cache := cache.NewSimpleCache()
	if cache == nil {
		t.Errorf("NewCache() should return a non-nil Cache instance")
	}

	// Test Set and Get methods
	cache.Set("key1", "value1")
	value, err := cache.Get("key1")
	if err != nil {
		t.Errorf("Get() should not return an error for existing keys")
	}
	if value != "value1" {
		t.Errorf("Get() should return the correct value for existing keys")
	}

	// Test Delete method
	cache.Delete("key1")
	value, err = cache.Get("key1")
	if err == nil {
		t.Errorf("Get() should return an error for deleted keys")
	}
	if value != "" {
		t.Errorf("Get() should return an empty string for deleted keys")
	}

}
