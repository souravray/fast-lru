package lru

import "errors"

// Base LRU implementation is not
// safe for concurent operations
type LRU struct {
	list    *List
	hmap    map[interface{}]*Node
	maxSize int
}

func newLRU(s int) (*LRU, error) {
	if s < 1 {
		return nil, errors.New("LRU size should be at least 1")
	}
	c := &LRU{
		newList(),
		make(map[interface{}]*Node, s),
		s,
	}
	return c, nil
}

// Len method returns the numbers of items in the cache
func (l *LRU) Len() int {
	return l.list.len
}

// Add method takes a key and value and add it to cache
func (l *LRU) Add(key, val interface{}) {
	// Check for existing item
	if n, ok := l.hmap[key]; ok {
		l.list.moveToFront(n)
		n.Value = val
		return
	}
	// Insert a new item
	n := l.list.insertAtFront(key, val)
	l.hmap[key] = n
	// Check size constrain
	if l.list.len > l.maxSize {
		en := l.list.removeFromBack()
		delete(l.hmap, en.Key)
	}
	return
}

// Exit method takes a key ane returns boolean value
// of cache hit/miss without updating the iteam sequence
func (l *LRU) Exist(key interface{}) bool {
	if _, ok := l.hmap[key]; ok {
		return true
	}
	return false
}

// Fetch method takes a key, move the item to front,
// and returns the value or nil and a boolean value
// of cache hit/miss
func (l *LRU) Fetch(key interface{}) (interface{}, bool) {
	if n, ok := l.hmap[key]; ok {
		n = l.list.moveToFront(n)
		if n == nil {
			return nil, false
		}
		return n.Value, true
	}
	return nil, false
}

// Keys method takes an order direction 'asc' or 'desc',
// and all keys in order of their recency. asc will return
// the most recent first, desc will return oldest first
func (l *LRU) Keys(odr string) ([]interface{}, error) {
	var li *ListIterator
	if odr == "asc" {
		li = l.list.Iterate("forward")
	} else if odr == "desc" {
		li = l.list.Iterate("backward")
	} else {
		return nil, errors.New("Unsupported order directive")
	}
	keys := make([]interface{}, l.Len())
	for li.Next() {
		i, n := li.Value()
		keys[i] = n.Key
	}
	return keys, nil
}

// Remove method deletes the provided key from the cache,
// returns boolean sucess/notf ound
func (l *LRU) Remove(key interface{}) bool {
	if n, ok := l.hmap[key]; ok {
		l.list.remove(n)
		delete(l.hmap, key)
		return true
	}
	return false
}
