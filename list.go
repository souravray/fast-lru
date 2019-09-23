package lru

/* An implementation of circular doubly linked
list since we don't want to use container/list
for the assignment ;)
*/
type list struct {
	root node // sentinel node
	len  int
}

// newList will create and initializ
// a doubly linked list instance
func newList() *list {
	l := &list{len: 0}
	l.root = node{}
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}

/* vaildation method */

// isValidMember method will return true,
// if the node belongs to the list and it
// is not the root node.
func (l *list) isValidMember(n *node) bool {
	if n.list != l || n == &l.root {
		return false
	}
	return true
}

/* getter setter methods */

// insertAtFront method will always insert at
// the start of the list - implemented for baseLRU
func (l *list) insertAtFront(key, val interface{}) *node {
	n := newNode(key, val)
	// take reference of cuurent
	// 1st node of list
	fn := l.root.next
	// insert node at root.next
	n.prev = &l.root
	n.next = fn
	fn.prev = n
	l.root.next = n
	n.list = l
	l.len++
	return n
}

// moveToFront method moves an exsiting
// node to the start of list.
func (l *list) moveToFront(n *node) *node {
	if !l.isValidMember(n) {
		return nil
	}
	// don't change anything, if the node
	// is already at the start of the list
	if n == l.root.next {
		return n
	}
	// take reference of current
	// 1st node of list
	fn := l.root.next
	// swap nodes
	n.prev.next = n.next
	n.next.prev = n.prev
	l.root.next = n
	n.next = fn
	n.prev = &l.root
	fn.prev = n
	return n
}

// remove method deletes an exsiting
// node from the list.
func (l *list) remove(n *node) *node {
	if !l.isValidMember(n) {
		return nil
	}
	// remove the node
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	n.list = nil
	l.len--
	return n
}

// removeFromBack method delete the last
// node of the list.
func (l *list) removeFromBack() *node {
	// take reference of current
	// last node of list
	ln := l.root.prev

	return l.remove(ln)
}

// Add a iterator interface for the list
type listIterator struct {
	direction string
	current   *node
	index     int
	list      *list
}

// iterate method takes a direction string
// and an listIterator instance
func (l *list) iterate(dir string) *listIterator {
	if dir != "forward" && dir != "backward" {
		dir = "forward"
	}
	li := &listIterator{
		direction: dir,
		list:      l,
		index:     -1,
		current:   &l.root,
	}
	return li
}

func (li *listIterator) value() (int, *node) {
	if li.current == &li.list.root {
		return li.index, nil
	}
	return li.index, li.current
}

func (li *listIterator) next() bool {
	li.index++
	if li.direction == "forward" {
		li.current = li.current.next
	} else {
		li.current = li.current.prev
	}
	if li.current == &li.list.root {
		return false
	}
	return true
}
