package piscine

func ListPushFront(l *List, data interface{}) {
	if l == nil {
		return
	}
	n := &NodeL{data, nil}
	if l.Head == nil {
		l.Tail = n
	} else {
		n.Next = l.Head
	}
	l.Head = n
}
