package main

import "fmt"

// EvictionAlgo is the interface for all eviction strategies
type EvictionAlgo interface {
	Evict()
}

// LRU is a concrete strategy that implements EvictionAlgo
type LRU struct{}

func (l *LRU) Evict() {
	fmt.Println("Evicting using LRU strategy")
}

// FIFO is a concrete strategy that implements EvictionAlgo
type FIFO struct{}

func (f *FIFO) Evict() {
	fmt.Println("Evicting using FIFO strategy")
}

// LFU is a concrete strategy that implements EvictionAlgo
type LFU struct{}

func (l *LFU) Evict() {
	fmt.Println("Evicting using LFU strategy")
}

// Cache is the main cache class that uses an eviction strategy
type Cache struct {
	maxSize  int
	strategy EvictionAlgo
	cache    map[string]interface{}
}

// NewCache is a constructor for the Cache
func NewCache(maxSize int, strategy EvictionAlgo) *Cache {
	return &Cache{
		maxSize:  maxSize,
		strategy: strategy,
		cache:    make(map[string]interface{}),
	}
}

// SetEvictionAlgo sets a new eviction strategy
func (c *Cache) SetEvictionAlgo(strategy EvictionAlgo) {
	c.strategy = strategy
}

// Evict triggers the eviction strategy
func (c *Cache) Evict() {
	c.strategy.Evict()
}

// Add adds a new item to the cache
func (c *Cache) Add(key string, value interface{}) {
	if len(c.cache) >= c.maxSize {
		c.Evict()
	}
	c.cache[key] = value
	fmt.Printf("Added %s: %v to cache\n", key, value)
}

// Get retrieves an item from the cache
func (c *Cache) Get(key string) interface{} {
	return c.cache[key]
}
