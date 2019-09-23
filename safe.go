package lru

import (
	"sync"
)

type LRU struct {
	lru  *baseLRU
	lock sync.RWMutex
}

func New(s int) (*LRU, error) {
	lru, err := newBaseLRU(s)
	if err != nil {
		return nil, err
	}
	slru := &LRU{lru: lru}
	return slru, nil
}

// Len method returns the numbers of items in the cache
func (sl *LRU) Len() int {
	sl.lock.Lock()
	ln := sl.lru.len()
	sl.lock.Unlock()
	return ln
}

// Add method takes a key and value and add it to cache
func (sl *LRU) Add(key, val interface{}) {
	sl.lock.Lock()
	sl.lru.add(key, val)
	sl.lock.Unlock()
}

// Exit method takes a key ane returns boolean value
// of cache hit/miss without updating the iteam sequence
func (sl *LRU) Exist(key interface{}) bool {
	sl.lock.Lock()
	ok := sl.lru.exist(key)
	sl.lock.Unlock()
	return ok
}

// Fetch method takes a key, move the item to front,
// and returns the value or nil and a boolean value
// of cache hit/miss
func (sl *LRU) Fetch(key interface{}) (interface{}, bool) {
	sl.lock.Lock()
	val, ok := sl.lru.fetch(key)
	sl.lock.Unlock()
	return val, ok
}

// Keys method takes an order direction 'asc' or 'desc',
// and all keys in order of their recency. asc will return
// the most recent first, desc will return oldest first
func (sl *LRU) Keys(odr string) ([]interface{}, error) {
	sl.lock.Lock()
	keys, err := sl.lru.keys(odr)
	sl.lock.Unlock()
	return keys, err
}

// Remove method deletes the provided key from the cache,
// returns boolean sucess/notf ound
func (sl *LRU) Remove(key interface{}) bool {
	sl.lock.Lock()
	ok := sl.lru.remove(key)
	sl.lock.Unlock()
	return ok
}
