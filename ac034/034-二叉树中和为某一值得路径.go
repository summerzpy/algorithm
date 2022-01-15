package ac034

/*
回溯
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	r := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		target -= root.Val
		if target == 0 && root.Left == nil && root.Right == nil {
			r = append(r, append([]int(nil),path...))
		}
		dfs(root.Left, target)
		dfs(root.Right, target)
		path = path[:len(path)-1]
	}
	dfs(root, target)
	return r
}
