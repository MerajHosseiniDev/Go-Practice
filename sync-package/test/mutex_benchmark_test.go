package test

import (
	"sync"
	"testing"
)

var value int = 42

func BenchmarkReadWithMutex(b *testing.B) {
	var mu sync.Mutex

	for i:=0; i< b.N; i++ {
		mu.Lock()
		_ = value
		mu.Unlock()
	}
}

func BenchmarkReadWithRWMutex(b *testing.B) {
	var mu sync.RWMutex

	for i:=0; i< b.N; i++ {
		mu.RLock()
		_ = value
		mu.RUnlock()
	}
}