package piscine

func ListClear(l *List) {
	if l == nil {
		return
	}
	link := l.Head
	for link != nil {
		link.Data = nil
		link, link.Next = link.Next, nil
	}
	l.Head = nil
	l.Tail = nil
	l = nil
}
