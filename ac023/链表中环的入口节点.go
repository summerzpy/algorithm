package ac023

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
注意公式推导
*/

func EntryNodeOfLoop(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return nil
	}
	fast := head
	slow := head
	for {
		if fast.Next == nil || fast.Next.Next == nil{
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
