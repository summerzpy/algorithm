package ac008

/*
题目：
	给定一棵二叉树和其中一个节点，如何找出中序遍历序列的下一个节点？树中的节点除了有两个分别指向
左、右子节点的指针，还有一个指向父节点的指针。
思路：
	主要判断右子树
*/

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

func GetNext(node *TreeLinkNode) *TreeLinkNode {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		next := node.Right
		for next.Left != nil {
			next = next.Left
		}
		return next
	} else {
		for node.Next != nil {
			if node.Next.Right == node{
				node = node.Next
			} else {
				return node.Next
			}
		}
	}
	return nil
}
