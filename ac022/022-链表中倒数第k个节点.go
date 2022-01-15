package ac022

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
快慢指针，注意两次边界条件的判断
*/

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	fast,slow := head,head
	for fast != nil && k>0 {
		fast = fast.Next
		k--
	}
	if k>0 {
		return nil
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
