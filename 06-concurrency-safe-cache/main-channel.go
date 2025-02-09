package main

import (
	"fmt"
)

type getRequest struct {
	key  string
	resp chan getResponse
}

type getResponse struct {
	value any
	ok    bool
}

type setRequest struct {
	key   string
	value any
	done  chan struct{}
}

type deleteRequest struct {
	key  string
	done chan struct{}
}

type ChanCache struct {
	getCh    chan getRequest
	setCh    chan setRequest
	deleteCh chan deleteRequest
	stopCh   chan struct{}
}

func NewChanCache() *ChanCache {
	c := &ChanCache{
		getCh:    make(chan getRequest),
		setCh:    make(chan setRequest),
		deleteCh: make(chan deleteRequest),
		stopCh:   make(chan struct{}),
	}
	go c.run()
	return c
}

func (c *ChanCache) run() {
	data := make(map[string]any)
	for {
		select {
		case req := <-c.getCh:
			value, ok := data[req.key]
			req.resp <- getResponse{value, ok}

		case req := <-c.setCh:
			data[req.key] = req.value
			req.done <- struct{}{}

		case req := <-c.deleteCh:
			delete(data, req.key)
			req.done <- struct{}{}

		case <-c.stopCh:
			close(c.getCh)
			close(c.setCh)
			close(c.deleteCh)
			return
		}
	}
}

func (c *ChanCache) Get(key string) (any, bool) {
	respCh := make(chan getResponse)
	c.getCh <- getRequest{
		key:  key,
		resp: respCh,
	}
	resp := <-respCh
	return resp.value, resp.ok
}

func (c *ChanCache) Set(key string, value any) {
	done := make(chan struct{})
	c.setCh <- setRequest{
		key:   key,
		value: value,
		done:  done,
	}
	<-done
}

func (c *ChanCache) Delete(key string) {
	done := make(chan struct{})
	c.deleteCh <- deleteRequest{
		key:  key,
		done: done,
	}
	<-done
}

func (c *ChanCache) Close() {
	close(c.stopCh)
}

func main() {
	chanCache := NewChanCache()

	chanCache.Set("name", "Alice")
	chanCache.Set("age", 30)

	if value, ok := chanCache.Get("name"); ok {
		fmt.Printf("name: %v\n", value)
	} else {
		fmt.Println("key 'name' not found")
	}

	if value, ok := chanCache.Get("age"); ok {
		fmt.Printf("age: %v\n", value)
	} else {
		fmt.Println("key 'age' not found")
	}

	if _, ok := chanCache.Get("nonexistent"); !ok {
		fmt.Println("key 'nonexistent' not found")
	}

	chanCache.Delete("name")
	if _, ok := chanCache.Get("name"); !ok {
		fmt.Println("key 'name' successfully deleted")
	}

	chanCache.Close()
}
