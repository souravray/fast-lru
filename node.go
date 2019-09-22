package lru

/* An implementation of linked-list node */

type Node struct {
	next, prev *Node

	// parent List
	list *List

	// The stored key (interface)
	Key interface{}

	// The stored value (interface)
	Value interface{}
}

func newNode(key, val interface{}) *Node {
	n := &Node{Key: key, Value: val}
	return n
}
