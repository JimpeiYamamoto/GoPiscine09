package piscine

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l == nil || comp == nil {
		return nil
	}
	it := l.Head
	for it != nil {
		if comp(it.Data, ref) == true {
			return &it.Data
		}
		it = it.Next
	}
	return nil
}
