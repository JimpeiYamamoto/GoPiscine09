package piscine

func ListReverse(l *List) {
	if l == nil {
		return
	}
	n := ListSize(l) - 1
	for n > 0 {
		ptr := l.Head
		for i := 0; i < n; i++ {
			ptr.Data, ptr.Next.Data = ptr.Next.Data, ptr.Data
			ptr = ptr.Next
		}
		n--
	}
}
