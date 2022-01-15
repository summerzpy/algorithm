package ac006

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
*/
func reversePrint(head *ListNode) []int {
	r := make([]int, 0)
	if head == nil {
		return r
	}
	return appendData(head)
}
//递归
func appendData(head *ListNode) []int{
	if head.Next != nil {
		list := appendData(head.Next)
		list = append(list, head.Val)
		return list
	}
	return []int{head.Val}
}
//如果限制链表长度的话，第二种方法时间复杂度最优
func printListFromTailToHead2( head *ListNode )[]int {
	if head == nil {
		return nil
	}
	r := make([]int,0)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	for i,j := 0,len(r)-1;i<j;{
		r[i],r[j] = r[j],r[i]
		i++
		j--
	}
	return r
}
