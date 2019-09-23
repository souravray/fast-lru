package lru

import "testing"

/* test helper methods */
func checkLength(t *testing.T, l *list, ln int) bool {
	if l.len != ln {
		t.Errorf("list length is %d, expected %d", l.len, ln)
		return false
	}
	return true
}

func checkConsistency(t *testing.T, l *list, ns []*node) {
	c := len(ns)
	if c == 0 {
		if l.root.next != &l.root || l.root.prev != &l.root {
			t.Errorf("Inconsistency found at root node (%p).prev = %p, (%p).next = %p", &l.root, l.root.prev, &l.root, l.root.next.prev)
		}
		return
	}
	pn := &l.root
	for i, n := range ns {
		if n.prev != pn {
			t.Errorf("Inconsistency at [%d]th node (%p).prev = %p, expected %p", i, n, n.prev, pn)
		}
		if pn.next != n {
			t.Errorf("Inconsistency at [%d]th node (%p).next = %p, expected %p", i, pn, n.next, n)
		}
		if i == c-1 {
			if n.next != &l.root || n != l.root.prev {
				t.Error("list is not circular")
			}
		}
		pn = n
	}
}

/*test cases*/
func Testlist(t *testing.T) {
	l := newList()
	// at zero node
	checkLength(t, l, 0)
	checkConsistency(t, l, []*node{})

	// with multiple nodes
	n1 := l.insertAtFront(nil, "Value 1")
	checkLength(t, l, 1)
	checkConsistency(t, l, []*node{n1})
	n2 := l.insertAtFront(nil, "Value 2")
	checkLength(t, l, 2)
	checkConsistency(t, l, []*node{n2, n1})
	n3 := l.insertAtFront(nil, "Value 3")
	checkLength(t, l, 3)
	checkConsistency(t, l, []*node{n3, n2, n1})
	n4 := l.insertAtFront(nil, "Value 4")
	checkLength(t, l, 4)
	checkConsistency(t, l, []*node{n4, n3, n2, n1})
}

func TestModifylist(t *testing.T) {
	l := newList()
	n1 := l.insertAtFront("Key1", "Value 1")
	n2 := l.insertAtFront("Key2", "Value 2")
	n3 := l.insertAtFront("Key3", "Value 3")
	n4 := l.insertAtFront("Key4", "Value 4")
	n5 := l.insertAtFront("Key5", "Value 5")
	checkLength(t, l, 5)
	checkConsistency(t, l, []*node{n5, n4, n3, n2, n1})

	l.moveToFront(n3)
	checkConsistency(t, l, []*node{n3, n5, n4, n2, n1})

	l.moveToFront(n1)
	checkConsistency(t, l, []*node{n1, n3, n5, n4, n2})

	l.remove(n1)
	checkLength(t, l, 4)
	checkConsistency(t, l, []*node{n3, n5, n4, n2})

	l.removeFromBack()
	checkLength(t, l, 3)
	checkConsistency(t, l, []*node{n3, n5, n4})

	l.removeFromBack()
	l.moveToFront(n4)
	checkLength(t, l, 2)
	checkConsistency(t, l, []*node{n3, n5})
}

func TestlistIterator(t *testing.T) {
	l := newList()
	n1 := l.insertAtFront("Key1", "Value 1")
	n2 := l.insertAtFront("Key2", "Value 2")
	n3 := l.insertAtFront("Key3", "Value 3")
	n4 := l.insertAtFront("Key4", "Value 4")
	n5 := l.insertAtFront("Key5", "Value 5")

	nArr := []*node{n5, n4, n3, n2, n1}
	fli := l.iterate("") // default should be forward
	for fli.next() {
		i, n := fli.value()
		if n != nArr[i] {
			t.Errorf("Iterator at %d returns node %p, expected %p", i, n, nArr[i])
		}
	}

	nRevArr := []*node{n1, n2, n3, n4, n5}
	rli := l.iterate("backward")
	for rli.next() {
		i, n := rli.value()
		if n != nRevArr[i] {
			t.Errorf("Iterator at %d returns node %p, expected %p", i, n, nRevArr[i])
		}
	}
}
