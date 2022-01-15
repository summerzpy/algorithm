package ac055

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(height(root.Left))-float64(height(root.Right))) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.Left),height(node.Right))
}
