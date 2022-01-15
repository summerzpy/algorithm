package ac024

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
需要设置两个指针 pre-记录当前节点的前一个节点；next-记录当前节点的后一个节点，防止链表断开
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	for head != nil {
		//保存当前节点的后续节点，防止链表断开
		next := head.Next
		//改变当前节点指针的指向，指向pre
		head.Next = pre
		//pre移动到当前节点
		pre = head
		//当前节点移动到后续节点
		head = next
	}
	return pre
}
