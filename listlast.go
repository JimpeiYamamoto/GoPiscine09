package piscine

func ListLast(l *List) interface{} {
	if l == nil {
		return nil
	}
	if l.Head == nil {
		return l.Head
	} else {
		it := l.Head
		for it.Next != nil {
			it = it.Next
		}
		return it.Data
	}
}
