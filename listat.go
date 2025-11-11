package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 || l == nil {
		return nil
	}

	current := l
	index := 0
	for current != nil {
		if index == pos {
			return current
		}
		current = current.Next
		index++
	}

	// If pos is greater than the number of nodes
	return nil
}
