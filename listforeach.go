package piscine

func ListForEach(l *List, f func(*NodeL)) {
	if l == nil || f == nil {
		return
	}
	it := l.Head
	for it != nil {
		f(it)
		it = it.Next
	}
}
