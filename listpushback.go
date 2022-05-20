package piscine

func ListPushBack(l *List, data interface{}) {
	if l == nil {
		return
	}
	node := &NodeL{data, nil}
	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
	}
	l.Tail = node
}
