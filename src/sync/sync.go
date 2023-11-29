package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// Mutexで排他ロック 並行処理されても各goroutineでの処理は順番を待つことになる
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

// constructのようなもの
func NewCounter() *Counter {
	return &Counter{}
}
