package ac028

/*
二叉树是否对称->左右子树是否镜像
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHandler(root.Left,root.Right)
}

func isSymmetricHandler(a *TreeNode,b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isSymmetricHandler(a.Right,b.Left) && isSymmetricHandler(a.Left,b.Right)
}