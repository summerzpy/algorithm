package ac036

func Convert2( root *TreeNode ) *TreeNode {
	if root == nil {
		return nil
	}
	var head, tail *TreeNode
	var f func(root *TreeNode)
	f = func (cur *TreeNode) {
		if cur ==nil {
			return
		}
		f(cur.Left)
		if tail == nil {
			head = cur
		} else {
			tail.Right = cur
			cur.Left = tail
		}
		tail = cur
		f(cur.Right)
	}
	f(root)
	//     head.Left = tail
	//     tail.Right = head
	return head
}
