package ac032

/*
题目：
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

思路：用栈，设置一个flag标识左右遍历顺序
*/

func levelOrder4(root *TreeNode) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}
	queue := []*TreeNode{root}
	flag := true
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, size)
		for i := 0;i<size;i++{
			t := queue[i]
			if flag {
				tmp[i] = t.Val
			} else {
				tmp[size-i-1] = t.Val
			}
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue =append(queue, t.Right)
			}
		}
		queue = queue[size:]
		flag = !flag
		r = append(r, tmp)
	}
	return r
}



