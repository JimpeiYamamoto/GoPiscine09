package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	link := l
	for i := 0; i < pos; i++ {
		if link == nil {
			return nil
		}
		link = link.Next
	}
	return link
}
