package ac035

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
/*
方法一：
第一次遍历原链表时不考虑random指针，只是复制节点；
第二次遍历原链表时只考虑random指针。
这种方法时间复杂度为O(n^2)，主要耗时在于确定新节点的random指针上。
方法二：
用空间换时间，使用map数据结构，保存原节点与新节点的对应关系，这样在确定random指针时可快速定位新节点。
时间复杂度为O(n)，空间复杂度为O(n)
方法三：
将新节点先插入至对应原节点的后面，最后将新节点一一拆分出来。
时间复杂度为O(n)，空间复杂度为O(1)
*/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		copyV := &Node{cur.Val, nil,nil}
		copyV.Next = cur.Next
		cur.Next = copyV
		cur = copyV.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	newHead := head.Next
	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = tmp.Next
		cur = tmp
	}
	return newHead
}
