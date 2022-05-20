package piscine

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	if l == nil || f == nil || cond == nil {
		return
	}
	it := l.Head
	for it != nil {
		if cond(it) {
			f(it)
		}
		it = it.Next
	}
}
