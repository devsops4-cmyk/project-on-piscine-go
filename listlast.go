package piscine

func ListLast(l *List) interface{} {
	if l.Tail == nil {
		// if no list found
		return nil
	}
	return l.Tail.Data // return l.tail.data
} // end of the program
