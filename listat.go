package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}
	count := 0
	for l != nil {
		if count == pos {
			return l
		}
		l = l.Next
		count++
	}
	return nil
}
