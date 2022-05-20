package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2 == nil {
		return
	}
	if l1 == nil {
		l1 = l2
		return
	}
	if l1.Head == nil {
		l1.Head, l1.Tail = l2.Head, l2.Tail
		return
	}
	it := l1.Head
	for it != nil && it.Next != nil {
		it = it.Next
	}
	it.Next = l2.Head
	for it.Next != nil {
		it = it.Next
	}
	l1.Tail = it
}
