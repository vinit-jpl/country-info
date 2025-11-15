package utils

import "sync"

type InMemoryCache struct {
	data map[string]interface{} // key: country name, value: API response
	mu   sync.RWMutex           // this prevents goroutines from reading/writing at the same time
}

func NewCache() Cache {

	// Create a Cache variable
	c := InMemoryCache{}

	// Initialize the map inside it
	c.data = make(map[string]interface{})

	// Return the address of this variable
	return &c
}

func (c *InMemoryCache) Get(key string) (interface{}, bool) {

	c.mu.RLock()           // lock the cache for reading, so that many goroutines can read at the same time
	defer c.mu.RUnlock()   // unlock after reading is done
	val, ok := c.data[key] // give the value and whether it was found

	return val, ok
}

func (c *InMemoryCache) Set(key string, value interface{}) {
	c.mu.Lock()         // lock the cache for writing
	defer c.mu.Unlock() // unlock after writing is done
	c.data[key] = value // store the value in the map
}
