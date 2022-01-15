package ac032

/*
题目：
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
*/
func levelOrder2(root *TreeNode) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmp := make([]int, 0)
		for i, l := 0,len(queue);i < l;i++ {
			t := queue[0]
			tmp = append(tmp, t.Val)
			queue = queue[1:]
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
		r = append(r, tmp)
	}
	return r
}
//写法二 巧用level变量
func levelOrder3(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	//利用队列
	queue := []*TreeNode{root}
	var level int = 0
	for len(queue) != 0 {
		//利用临时队列，暂存每个节点的左右子树
		temp := []*TreeNode{}
		//只需考虑在同一层上追加元素
		res = append(res,[]int{})
		//遍历队列，追加队列元素到切片同一层，追加队列元素左右子树到临时队列
		for _,v := range queue {
			res[level] = append(res[level],v.Val)
			if v.Left != nil {
				temp = append(temp,v.Left)
			}
			if v.Right != nil{
				temp = append(temp,v.Right)
			}
		}
		//层级加1，队列重新复制为队列的左右子树集
		level++
		//队列赋值
		queue = temp
	}
	return res
}

//levelOrder2优化写法
func levelOrder5(root *TreeNode) [][]int {
	r := make([][]int, 0)
	if root == nil {
		return r
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < l;i++ {
			t := queue[i]
			tmp = append(tmp, t.Val)
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
		queue = queue[l:]
		r = append(r, tmp)
	}
	return r
}

