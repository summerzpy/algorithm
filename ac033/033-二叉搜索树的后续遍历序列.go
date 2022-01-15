package ac033

/*
按照二叉搜索树的特征 循环判断
设置i,j指针，表示左右边界
*/


func verifyPostorder(postorder []int) bool {
	var verify func(i, j int) bool
	verify = func(i,j int) bool {
		for i >= j {
			return true
		}
		index := i
		for postorder[index] < postorder[j] {
			index++
		}
		right := index
		for postorder[index] > postorder[j] {
			index++
		}
		return index == j && verify(i,right-1) && verify(right, j-1)
	}
	return verify(0,len(postorder)-1)
}