package ac054

/*
第k大：右 根 左
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret *TreeNode
var cnt int

func KthNode(root *TreeNode, k int) *TreeNode {
	inOrder(root, k)
	return ret
}

func inOrder(node *TreeNode, k int) {
	if node == nil || cnt >= k {
		return
	}
	inOrder(node.Left, k)
	cnt++
	if cnt == k {
		ret = node
	}
	inOrder(node.Right, k)
}
