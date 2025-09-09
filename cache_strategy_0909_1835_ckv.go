// 代码生成时间: 2025-09-09 18:35:11
package main

import (
    "fmt"
    "log"
    "time"
)

// CacheItem defines the structure for a cache item.
type CacheItem struct {
    Value    string    // The cached value.
    Expiry   time.Time // The time at which the item expires.
}

// Cache provides an interface for cache operations.
type Cache interface {
    Get(key string) (*CacheItem, error)
    Set(key string, item *CacheItem) error
    Invalidate(key string) error
}

// SimpleCache is a simple in-memory cache implementation.
type SimpleCache struct {
    store  map[string]CacheItem
    expiry time.Duration
}

// NewSimpleCache creates a new SimpleCache with a given expiry duration.
func NewSimpleCache(expiry time.Duration) *SimpleCache {
    return &SimpleCache{store: make(map[string]CacheItem), expiry: expiry}
}

// Get retrieves a cached item by key.
func (c *SimpleCache) Get(key string) (*CacheItem, error) {
    item, exists := c.store[key]
    if !exists {
        return nil, fmt.Errorf("item not found")
    }
    if time.Now().After(item.Expiry) {
        // Item has expired, remove it from the cache.
        delete(c.store, key)
        return nil, fmt.Errorf("item expired")
    }
    return &item, nil
}

// Set adds or updates a cached item.
func (c *SimpleCache) Set(key string, item *CacheItem) error {
    if item == nil {
        return fmt.Errorf("cannot set a nil item")
    }
    item.Expiry = time.Now().Add(c.expiry)
    c.store[key] = *item
    return nil
}

// Invalidate removes a cached item by key.
func (c *SimpleCache) Invalidate(key string) error {
    _, exists := c.store[key]
    if !exists {
        return fmt.Errorf("item not found")
    }
    delete(c.store, key)
    return nil
}

func main() {
    cache := NewSimpleCache(5 * time.Minute) // Cache expiry set to 5 minutes.

    // Set a value in the cache.
    if err := cache.Set("key1", &CacheItem{Value: "value1"}); err != nil {
        log.Fatal(err)
    }

    // Get the value from the cache.
    if item, err := cache.Get("key1"); err == nil {
        fmt.Println("Cached value: ", item.Value)
    } else {
        log.Fatal(err)
    }

    // Invalidate the cached value.
    if err := cache.Invalidate("key1"); err != nil {
        log.Fatal(err)
    }
}
