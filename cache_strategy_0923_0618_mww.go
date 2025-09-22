// 代码生成时间: 2025-09-23 06:18:55
package main

import (
    "fmt"
    "sync"
    "time"
)

// CacheEntry represents a cached entry with a value and an expiration time.
type CacheEntry struct {
    Value    interface{}
    Expiry   time.Time
}

// Cache is a simple in-memory cache with expiration.
type Cache struct {
    entries map[string]*CacheEntry
    lock    sync.RWMutex
}

// NewCache creates a new Cache instance.
func NewCache() *Cache {
    return &Cache{
        entries: make(map[string]*CacheEntry),
    }
}

// Set adds or updates a cache entry with an expiration duration.
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
    c.lock.Lock()
    defer c.lock.Unlock()
    entry := &CacheEntry{
        Value:    value,
        Expiry:   time.Now().Add(duration),
    }
    c.entries[key] = entry
}

// Get retrieves a value from the cache. If the entry is expired or not found, it returns nil.
func (c *Cache) Get(key string) (interface{}, error) {
    c.lock.RLock()
    defer c.lock.RUnlock()
    entry, ok := c.entries[key]
    if !ok || entry.Expiry.Before(time.Now()) {
        return nil, fmt.Errorf("entry not found or expired")
    }
    return entry.Value, nil
}

// Evict removes an entry from the cache.
func (c *Cache) Evict(key string) {
    c.lock.Lock()
    defer c.lock.Unlock()
    delete(c.entries, key)
}

// Clear removes all entries from the cache.
func (c *Cache) Clear() {
    c.lock.Lock()
    defer c.lock.Unlock()
    for key := range c.entries {
        delete(c.entries, key)
    }
}

func main() {
    // Create a new cache instance.
    cache := NewCache()

    // Set a value with a 5-second expiration.
    cache.Set("example", 42, 5*time.Second)

    // Retrieve the value.
    value, err := cache.Get("example")
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Retrieved value: %v
", value)
    }

    // Wait for 6 seconds to let the cache entry expire.
    time.Sleep(6 * time.Second)

    // Try to retrieve the expired value.
    value, err = cache.Get("example")
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("Retrieved value: %v
", value)
    }
}