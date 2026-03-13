package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	var last interface{}
	for node := l.Head; node != nil; node = node.Next {
		last = node.Data
	}
	return last
}
