package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}
	it := l.Head
	for it != nil && it.Data == data_ref {
		it, it.Next, it.Data = it.Next, nil, nil
	}
	l.Head = it
	it = l.Head
	for it != nil {
		for it.Next != nil && it.Next.Data == data_ref {
			it.Next.Data, it.Next = nil, it.Next.Next
		}
		if it.Next == nil {
			l.Tail = it
		}
		it = it.Next
	}
}
