package ac026

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
题目：
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。
注意B为空子树的判断
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isSubStructureHandler(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
}

func isSubStructureHandler(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return isSubStructureHandler(A.Left,B.Left) && isSubStructureHandler(A.Right,B.Right)
}
