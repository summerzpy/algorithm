package ac036

var pre *TreeNode //必须在全局变量上才可以实现

//递归
func Convert1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	helper(root)
	head, tail := root, root
	for head.Left != nil {
		head = head.Left
	}
	for tail.Right != nil {
		tail = tail.Right
	}
	head.Left = tail
	tail.Right = head
	return head
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left)
	if pre != nil {
		root.Left = pre
		pre.Right = root
	}
	pre = root
	helper(root.Right)
}
