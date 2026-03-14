package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	var before *NodeL = nil
	head := l.Head
	l.Tail = l.Head

	for head != nil {
		next := head.Next
		head.Next = before
		before = head
		head = next
	}

	l.Head = before
}
