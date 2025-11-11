package piscine

func ListClear(l *List) {
	// Iterate through the list
	current := l.Head // current head
	for current != nil {
		next := current.Next
		current.Next = nil
		current.Data = nil
		current = next
	}
	l.Head = nil
	l.Tail = nil
} // end of the program
