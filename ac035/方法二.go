package ac035


func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	p := head.Next
	pre := &newHead
	connection := make(map[*Node]*Node)
	connection[head] = pre
	for p != nil {
		newNode := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		pre.Next = newNode
		connection[p] = newNode
		p = p.Next
		pre = pre.Next
	}
	p = head
	newP := &newHead
	for p != nil {
		if p.Random != nil {
			newP.Random = connection[p.Random]
		}
		p = p.Next
		newP = newP.Next
	}
	return &newHead
}
