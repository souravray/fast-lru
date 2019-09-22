package lru

import (
  "testing"
  "fmt"
)
func checkLength(t *testing.T, l *List, len int) bool {
  if l.len != len {
    t.Errorf("List length is %d, expected %d", l.len, len)
    return false
  }
  return true
}

func checkConsistency(t *testing.T, l *List, ns []*Node) {
  c := len(ns)
  if c == 0 {
    if l.root.next != &l.root || l.root.prev != &l.root {
      t.Errorf("Inconsistency found at root node (%p).prev = %p, (%p).next = %p", &l.root, l.root.prev, &l.root, l.root.next.prev)
    }
    return
  }
  pn := &l.root
  fmt.Printf("\n\n\t%p  p-> %p n-> %p\n", &l.root, l.root.prev, l.root.next)
  for i, n := range ns {
    fmt.Printf("\t%p  p-> %p n-> %p\n", n, n.prev, n.next)
    if n.prev != pn {
      t.Errorf("Inconsistency at [%d]th node (%p).prev = %p, expected %p", i, n, n.prev, pn)
    }
    if pn.next != n {
      t.Errorf("Inconsistency at [%d]th node (%p).next = %p, expected %p", i, pn, n.next, n)
    }
    if i == c-1 {
      if n.next != &l.root || n != l.root.prev {
        t.Error("List is nit circular")
      }
    }
    pn = n
  }
}

func TestList(t *testing.T) {
  l := newList()
  // at zero node
  checkLength(t, l, 0)
  checkConsistency(t, l, []*Node{})
  
  // with multiple nodes
  n1 := l.insertAtFront("Value 1")
  checkLength(t, l, 1)
  checkConsistency(t, l, []*Node{ n1})
  n2 := l.insertAtFront("Value 2")
  checkLength(t, l, 2)
  checkConsistency(t, l, []*Node{ n2, n1})
  n3 := l.insertAtFront("Value 3")
  checkLength(t, l, 3)
  checkConsistency(t, l, []*Node{ n3, n2, n1})
  n4 := l.insertAtFront("Value 4")
  checkLength(t, l, 4)
  checkConsistency(t, l, []*Node{ n4, n3, n2, n1})
}

func TestModifyList(t *testing.T) {
  l := newList()
  n1 := l.insertAtFront("Value 1")
  n2 := l.insertAtFront("Value 2")
  n3 := l.insertAtFront("Value 3")
  n4 := l.insertAtFront("Value 4")
  n5 := l.insertAtFront("Value 5")
  checkLength(t, l, 5)
  checkConsistency(t, l, []*Node{ n5, n4, n3, n2, n1})

  l.moveToFront(n3)
  checkConsistency(t, l, []*Node{ n3, n5, n4, n2, n1})

  l.moveToFront(n1)
  checkConsistency(t, l, []*Node{ n1, n3, n5, n4, n2})

  l.remove(n1)
  checkLength(t, l, 4)
  checkConsistency(t, l, []*Node{ n3, n5, n4, n2})

  l.removeFromBack()
  checkLength(t, l, 3)
  checkConsistency(t, l, []*Node{ n3, n5, n4})

  l.removeFromBack()
  l.moveToFront(n4)
  checkLength(t, l, 2)
  checkConsistency(t, l, []*Node{ n3, n5})
}