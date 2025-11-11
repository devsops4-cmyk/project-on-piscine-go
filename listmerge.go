package piscine

// Write a function ListMerge that places elements of a list l2 at the end of another list
func ListMerge(l1 *List, l2 *List) {
	if l2.Head == nil {
		return
	}

	// If l1 is empty, l1 becomes l2
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// Link l1's tail to l2's head
	l1.Tail.Next = l2.Head

	// Update l1's tail to l2's tail
	l1.Tail = l2.Tail
} // end of the program
