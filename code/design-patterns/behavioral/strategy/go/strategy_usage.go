package main

func main() {
	lru := &LRU{}
	fifo := &FIFO{}
	cache := NewCache(2, lru)

	cache.Add("a", 1)
	cache.Add("b", 2)
	cache.Add("c", 3) // This should trigger eviction with LRU

	cache.SetEvictionAlgo(fifo)
	cache.Add("d", 4) // This should trigger eviction with FIFO
}
