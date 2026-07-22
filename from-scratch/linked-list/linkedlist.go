package singleLinkedList

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
}

func (l *List) Insert(v int) {
	newNode := &Node{Value: v}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (l *List) Delete(v int) bool {
	curr := l.Head
	prew := l.Head
	for curr != nil && curr.Value != v {
		prew = curr
		curr = curr.Next
	}
	if curr == nil {
		return false
	}
	if curr == l.Head {
		l.Head = curr.Next
		return true
	}
	prew.Next = curr.Next

	return true
}

func (l *List) Find(v int) *Node {
	curr := l.Head
	for curr != nil && curr.Value != v {
		curr = curr.Next
	}
	return curr
}

func (l *List) Middle() *Node {
	curr := l.Head
	counter := 0
	for curr != nil {
		counter += 1
		curr = curr.Next
	}
	counter = counter / 2

	curr = l.Head
	for range counter {
		curr = curr.Next
	}

	return curr
}

func (l *List) Reverse() {
	var prev *Node
	curr := l.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	l.Head = prev
}
func (l *List) HasCycle() bool {
	fast := l.Head
	slow := l.Head
	counter := 0
	for fast != nil {
		fast = fast.Next
		counter += 1
		if counter%2 == 0 {
			slow = slow.Next
		}

		if fast == slow {
			return true
		}
	}
	return false
}
