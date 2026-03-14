package piscine

func IsPositiveNode(node *NodeL) bool {
	if v, ok := node.Data.(int); ok {
		return v > 0
	}
	if v, ok := node.Data.(float32); ok {
		return v > 0
	}
	if v, ok := node.Data.(float64); ok {
		return v > 0
	}
	if v, ok := node.Data.(byte); ok {
		return v > 0
	}
	return false
}

func IsAlNode(node *NodeL) bool {
	if _, ok := node.Data.(int); ok {
		return false
	}
	if _, ok := node.Data.(float32); ok {
		return false
	}
	if _, ok := node.Data.(float64); ok {
		return false
	}
	if _, ok := node.Data.(byte); ok {
		return false
	}
	return true
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	head := l.Head
	for head != nil {
		if cond(head) {
			f(head)
		}
		head = head.Next
	}
}
