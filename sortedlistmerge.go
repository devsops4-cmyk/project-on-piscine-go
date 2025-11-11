package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	// Case 1: list is empty or new node goes before head
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	// Traverse until we find the correct position
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	// Insert newNode after current
	newNode.Next = current.Next
	current.Next = newNode

	return l
} // END OF MY PROGRAM
