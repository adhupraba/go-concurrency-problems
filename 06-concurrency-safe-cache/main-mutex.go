package main

import (
	"fmt"
	"sync"
)

type MuCache struct {
	mu   sync.RWMutex
	data map[string]any
}

func (this *MuCache) Get(key string) (val any, ok bool) {
	this.mu.RLock()
	defer this.mu.RUnlock()

	if val, ok := this.data[key]; ok {
		return val, ok
	}

	return nil, false
}

func (this *MuCache) Set(key string, value any) {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.data[key] = value
}

func (this *MuCache) Delete(key string) {
	this.mu.Lock()
	defer this.mu.Unlock()

	delete(this.data, key)
}

func NewMuCache() *MuCache {
	return &MuCache{
		mu:   sync.RWMutex{},
		data: make(map[string]any),
	}
}

func main() {
	muCache := NewMuCache()

	muCache.Set("name", "Alice")
	muCache.Set("age", 30)

	var wg sync.WaitGroup
	keys := []string{"name", "age", "nonexistent"}

	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			if val, ok := muCache.Get(k); ok {
				fmt.Printf("Key: %s, Value: %v\n", k, val)
			} else {
				fmt.Printf("Key: %s not found\n", k)
			}
		}(key)
	}

	wg.Wait()

	muCache.Delete("name")

	if _, ok := muCache.Get("name"); !ok {
		fmt.Println("Key 'name' successfully deleted")
	}
}
