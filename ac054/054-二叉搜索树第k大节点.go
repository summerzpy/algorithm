package ac054

var skip int
var res int
func kthLargest(root *TreeNode, k int) int {
	skip = k
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode){
	if root != nil{
		dfs(root.Right)
		skip--
		if skip == 0{
			res = root.Val
			return
		}
		dfs(root.Left)
	}
}

func kthLargest1(root *TreeNode, k int) int {
	var find func(root *TreeNode)
	res := 0
	find = func(cur *TreeNode) {
		if cur == nil || k <= 0{
			return
		}
		find(cur.Right)
		k--
		if k == 0 {
			res = cur.Val
			return
		}
		find(cur.Left)
	}
	find(root)
	return res
}
