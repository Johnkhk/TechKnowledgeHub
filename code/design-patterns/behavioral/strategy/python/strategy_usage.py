from cache import Cache, FIFO, LRU

cache = Cache(2, LRU())
cache.add("a", 1)
cache.add("b", 2)
cache.add("c", 3)  # This should trigger eviction
cache.set_eviction_algo(FIFO())
cache.add("d", 4)  # This should trigger eviction using FIFO strategy
