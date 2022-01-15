package ac055

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树的深度
func TreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(TreeDepth(root.Right), TreeDepth(root.Left))
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
