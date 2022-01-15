package ac007

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
边界条件判断
preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) || len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	treeIndex := 0
	for i := 0; i < len(preorder); i++ {
		if inorder[i] == rootVal {
			treeIndex = i
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:treeIndex+1], inorder[0:treeIndex]),
		Right: buildTree(preorder[treeIndex+1:], inorder[treeIndex+1:]),
	}
}
