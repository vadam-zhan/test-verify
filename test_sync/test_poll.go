package test_sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	// Test cases for pool functionality
}

// 使用 benchmark 测试，withPool 的性能会显著优于 withoutPool

// 使用 sync.Pool 复用字节切片
func withPool() {
	pool := &sync.Pool{
		New: func() interface{} { return make([]byte, 1024) },
	}
	for i := 0; i < 10000; i++ {
		b := pool.Get().([]byte)
		// 使用 b...
		pool.Put(b)
	}
}

// 每次都新建字节切片
func withoutPool() {
	for i := 0; i < 10000; i++ {
		b := make([]byte, 1024)
		// 使用 b...
		// b 在此作用域结束后被垃圾回收
		fmt.Println(b)
	}
}
