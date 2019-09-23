package lru

/* An implementation of linked-list node */

type node struct {
	next, prev *node

	// parent list
	list *list

	// The stored key (interface)
	key interface{}

	// The stored value (interface)
	value interface{}
}

func newNode(key, val interface{}) *node {
	n := &node{key: key, value: val}
	return n
}
