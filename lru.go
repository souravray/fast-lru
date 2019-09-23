package lru

import "errors"

// Base baseLRU implementation is not
// safe for concurent operations
type baseLRU struct {
	list    *list
	hmap    map[interface{}]*node
	maxSize int
}

func newBaseLRU(s int) (*baseLRU, error) {
	if s < 1 {
		return nil, errors.New("baseLRU size should be at least 1")
	}
	c := &baseLRU{
		newList(),
		make(map[interface{}]*node, s),
		s,
	}
	return c, nil
}

// len method returns the numbers of items in the cache
func (l *baseLRU) len() int {
	return l.list.len
}

// add method takes a key and value and add it to cache
func (l *baseLRU) add(key, val interface{}) {
	// Check for existing item
	if n, ok := l.hmap[key]; ok {
		l.list.moveToFront(n)
		n.value = val
		return
	}
	// Insert a new item
	n := l.list.insertAtFront(key, val)
	l.hmap[key] = n
	// Check size constrain
	if l.list.len > l.maxSize {
		en := l.list.removeFromBack()
		delete(l.hmap, en.key)
	}
	return
}

// exit method takes a key ane returns boolean value
// of cache hit/miss without updating the iteam sequence
func (l *baseLRU) exist(key interface{}) bool {
	if _, ok := l.hmap[key]; ok {
		return true
	}
	return false
}

// fetch method takes a key, move the item to front,
// and returns the value or nil and a boolean value
// of cache hit/miss
func (l *baseLRU) fetch(key interface{}) (interface{}, bool) {
	if n, ok := l.hmap[key]; ok {
		n = l.list.moveToFront(n)
		if n == nil {
			return nil, false
		}
		return n.value, true
	}
	return nil, false
}

// keys method takes an order direction 'asc' or 'desc',
// and all keys in order of their recency. asc will return
// the most recent first, desc will return oldest first
func (l *baseLRU) keys(odr string) ([]interface{}, error) {
	var li *listIterator
	if odr == "asc" {
		li = l.list.iterate("forward")
	} else if odr == "desc" {
		li = l.list.iterate("backward")
	} else {
		return nil, errors.New("Unsupported order directive")
	}
	keys := make([]interface{}, l.len())
	for li.next() {
		i, n := li.value()
		keys[i] = n.key
	}
	return keys, nil
}

// remove method deletes the provided key from the cache,
// returns boolean sucess/notf ound
func (l *baseLRU) remove(key interface{}) bool {
	if n, ok := l.hmap[key]; ok {
		l.list.remove(n)
		delete(l.hmap, key)
		return true
	}
	return false
}
