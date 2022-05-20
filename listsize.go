package piscine

func ListSize(l *List) int {
	if l == nil {
		return -1
	}
	var ret int
	it := l.Head
	for it != nil {
		it = it.Next
		ret++
	}
	return ret
}
