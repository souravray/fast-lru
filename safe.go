package lru

import (
	"sync"
)

type SafeLRU struct {
	lru  *LRU
	lock sync.RWMutex
}

func New(s int) (*SafeLRU, error) {
	lru, err := newLRU(s)
	if err != nil {
		return nil, err
	}
	slru := &SafeLRU{lru: lru}
	return slru, nil
}

// Len method returns the numbers of items in the cache
func (sl *SafeLRU) Len() int {
	sl.lock.Lock()
	ln := sl.lru.Len()
	sl.lock.Unlock()
	return ln
}

// Add method takes a key and value and add it to cache
func (sl *SafeLRU) Add(key, val interface{}) {
	sl.lock.Lock()
	sl.lru.Add(key, val)
	sl.lock.Unlock()
}

// Exit method takes a key ane returns boolean value
// of cache hit/miss without updating the iteam sequence
func (sl *SafeLRU) Exist(key interface{}) bool {
	sl.lock.Lock()
	ok := sl.lru.Exist(key)
	sl.lock.Unlock()
	return ok
}

// Fetch method takes a key, move the item to front,
// and returns the value or nil and a boolean value
// of cache hit/miss
func (sl *SafeLRU) Fetch(key interface{}) (interface{}, bool) {
	sl.lock.Lock()
	val, ok := sl.lru.Fetch(key)
	sl.lock.Unlock()
	return val, ok
}

// Keys method takes an order direction 'asc' or 'desc',
// and all keys in order of their recency. asc will return
// the most recent first, desc will return oldest first
func (sl *SafeLRU) Keys(odr string) ([]interface{}, error) {
	sl.lock.Lock()
	keys, err := sl.lru.Keys(odr)
	sl.lock.Unlock()
	return keys, err
}

// Remove method deletes the provided key from the cache,
// returns boolean sucess/notf ound
func (sl *SafeLRU) Remove(key interface{}) bool {
	sl.lock.Lock()
	ok := sl.lru.Remove(key)
	sl.lock.Unlock()
	return ok
}
