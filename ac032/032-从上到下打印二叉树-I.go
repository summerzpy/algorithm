package ac032

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
题目：
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
*/
func levelOrder(root *TreeNode) []int {
	r := make([]int,0)
	queue := make([]*TreeNode, 0)
	if root == nil {
		return r
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		for i:= 0;i<len(queue);i++ {
			t := queue[0]
			r = append(r, t.Val)
			queue = queue[1:]
			if t.Left != nil {
				queue = append(queue,t.Left)
			}
			if t.Right != nil {
				queue = append(queue,t.Right)
			}
		}
	}
	return r
}
