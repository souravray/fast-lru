package lru

import (
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkCacheAtFixedFreq(b *testing.B) {
	c, _ := newBaseLRU(10000)
	payloads := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		if i%2 == 0 {
			payloads[i] = rand.Int63() % 20000
		} else {
			payloads[i] = rand.Int63() % 40000
		}
	}
	b.ResetTimer()
	var hit, miss int
	for i := 0; i < 2*b.N; i++ {
		if i%2 == 0 {
			c.add(payloads[i], payloads[i])
		} else {
			if _, ok := c.fetch(payloads[i]); ok {
				hit++
			} else {
				miss++
			}
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func BenchmarkSafeCacheAtFixedFreq(b *testing.B) {
	c, _ := New(10000)
	payloads := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		if i%2 == 0 {
			payloads[i] = rand.Int63() % 20000
		} else {
			payloads[i] = rand.Int63() % 40000
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Add(payloads[i], payloads[i])
	}
	var hit, miss int
	for i := 0; i < b.N; i++ {
		if _, ok := c.Fetch(payloads[i]); ok {
			hit++
		} else {
			miss++
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func BenchmarkSafeCacheAtRandFreq(b *testing.B) {
	c, _ := New(10000)
	payloads := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		payloads[i] = rand.Int63() % 40000
	}
	b.ResetTimer()
	var hit, miss int
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			c.Add(payloads[i], payloads[i])
		} else {
			_, ok := c.Fetch(payloads[i])
			if ok {
				hit++
			} else {
				miss++
			}
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func BenchmarkConcurrentSafeCacheAtFixedFreq(b *testing.B) {
	c, _ := New(10000)
	payloads := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		if i%2 == 0 {
			payloads[i] = rand.Int63() % 20000
		} else {
			payloads[i] = rand.Int63() % 40000
		}
	}
	b.ResetTimer()
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(key int64) {
			c.Add(key, key)
			wg.Done()
		}(payloads[i])
	}
	wg.Wait()
	var h, m int
	for i := 0; i < b.N; i++ {
		go func(key int64) {
			if _, ok := c.Fetch(key); ok {
				h++
			} else {
				m++
			}
		}(payloads[i])
	}
	b.Logf("hit: %d miss: %d ratio: %f", h, m, float64(h)/float64(m))
}

func BenchmarkConcurrentSafeCacheAtRandFreq(b *testing.B) {
	c, _ := New(10000)
	payloads := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		payloads[i] = rand.Int63() % 40000
	}
	b.ResetTimer()
	var hit, miss int
	for i := 0; i < b.N; i++ {
			go func(key int64) {
			if i%2 == 0 {
					_, ok := c.Fetch(key)
					if ok {
						hit++
					} else {
						miss++
					}
			} else {
				
					c.Add(key, key)
				}
			}(payloads[i])
		}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}
