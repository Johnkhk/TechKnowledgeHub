from abc import ABC, abstractmethod


class EvictionAlgo(ABC):
    @abstractmethod
    def evict(self):
        pass


class LRU(EvictionAlgo):
    def evict(self):
        print("Evicting using LRU strategy")


class FIFO(EvictionAlgo):
    def evict(self):
        print("Evicting using FIFO strategy")


class LFU(EvictionAlgo):
    def evict(self):
        print("Evicting using LFU strategy")


class Cache:
    def __init__(self, max_size, strategy: EvictionAlgo):
        self.max_size = max_size
        self.strategy = strategy
        self.cache = {}

    def evict(self):
        self.strategy.evict()

    def set_eviction_algo(self, strategy: EvictionAlgo):
        self.strategy = strategy

    def add(self, key, value):
        if len(self.cache) >= self.max_size:
            self.evict()
        self.cache[key] = value
        print(f"Added {key}: {value} to cache")

    def get(self, key):
        return self.cache.get(key, None)
