package piscine

func ListReverse(l *List) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	var prev *NodeL = nil
	current := l.Head
	var next *NodeL

	l.Tail = l.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}
