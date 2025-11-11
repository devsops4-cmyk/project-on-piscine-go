package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		// Empty list: both Head and Tail point to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Non-empty list: insert at the beginning
		newNode.Next = l.Head
		l.Head = newNode
	}
}
