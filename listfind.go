package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

// program that returns the address of the data interface of the first node in the list
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head    // current head
	for current != nil { // when current is not equal to nil
		if comp(current.Data, ref) {
			return &current.Data
		}
		current = current.Next
	}
	return nil
} // end of the program end
