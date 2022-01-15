package ac036


/*
中序遍历
*/


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func Convert(root *TreeNode ) *TreeNode {
	if root == nil {
		return nil
	}
	st := []*TreeNode{}
	var pre, head *TreeNode
	for len(st) > 0 || root != nil {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}
		cur := st[len(st)-1]
		st = st[:len(st)-1]
		if pre != nil {
			pre.Right = cur
		} else {
			head = cur
		}
		cur.Left = pre
		pre = cur
		root = cur.Right
	}
	// 牛客不要求循环链表
	//     head.Left = pre
	//     pre.Right = head
	return head
}
