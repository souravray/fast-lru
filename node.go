package lru

/* An implementation of linked-list node */

type Node struct {
  next, prev *Node

  // parent List
  list *List

  // The stored value (interface)
  Value interface{}
}

func newNode(val interface{}) *Node {
  n := &Node{Value:val}
  return n
}
