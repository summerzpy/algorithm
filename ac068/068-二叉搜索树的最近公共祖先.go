package ac068

//二叉查找树 又名二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//迭代：时间复杂度 O(N)： 其中 NN 为二叉树节点数；每循环一轮排除一层，二叉搜索树的层数最小为 \log NlogN （满二叉树），最大为 NN （退化为链表）。
//空间复杂度：O(1)
func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	for root != nil {
		if root.Val < p && root.Val < q {
			root = root.Right
		} else if root.Val > p && root.Val > q {
			root = root.Left
		} else {
			break
		}
	}
	return root.Val
}

//递归 时间复杂度O(n)
//空间复杂度O(n)
func lowestCommonAncestor2(root *TreeNode, p int, q int) int {
	if root.Val < p && root.Val < q {
		return lowestCommonAncestor2(root.Right, p, q)
	}
	if root.Val > p && root.Val > q {
		return lowestCommonAncestor2(root.Left, p, q)
	}
	return root.Val
}
