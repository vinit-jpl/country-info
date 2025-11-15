package utils

import (
	"sync"
	"testing"
)

func TestInMemoryCacheRace(t *testing.T) {
	cache := NewCache()

	var wg sync.WaitGroup

	// run 1000 concurrent writers
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set("key", i)
		}(i)
	}

	// run 1000 concurrent readers
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Get("key")
		}()
	}

	wg.Wait()
}
