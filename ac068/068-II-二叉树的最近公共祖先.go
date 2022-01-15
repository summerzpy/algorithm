package ac068


//递归
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	//可优化为
	//if root == nil || root.Val == p.Val || root.Val == q.Val {
	//	return root
	//}
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left == nil {
		return right
	} else {
		return left
	}
}
