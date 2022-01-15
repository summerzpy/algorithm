package ac052

//map存放其中一个，另一个遍历对比
//双栈倒着比
//先找到两链表长度相等的地方，在遍历比较 时间复杂度O(m+n),空间复杂度O(1)
//设 A 的长度为 a + c，B 的长度为 b + c，其中 c 为尾部公共部分长度，可知 a + c + b = b + c + a。
//
//***当访问链表 A 的指针访问到链表尾部时，令它从链表 B 的头部重新开始访问链表 B；
//同样地，当访问链表 B 的指针访问到链表尾部时，
//令它从链表 A 的头部重新开始访问链表 A。
//这样就能控制访问 A 和 B 两个链表的指针能同时访问到交点。

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lengthA, lengthB := 0, 0
	for curA != nil {
		lengthA++
		curA = curA.Next
	}
	for curB != nil {
		lengthB++
		curB = curB.Next
	}
	var step int
	var fast, slow *ListNode
	if lengthA > lengthB {
		step = lengthA - lengthB
		fast, slow = headA, headB
	} else {
		step = lengthB - lengthA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func GetIntersectionNode2(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
