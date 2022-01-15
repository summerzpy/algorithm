package ac031

/*
借助辅助栈
1、借助一个辅助栈，模拟一遍弹出的操作
2、每插入一个元素到辅助栈中，判断是否可进行弹出操作
3、直到弹出全部元素，如果能够弹出，说明成功，否则失败。
*/

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack) - 1] == popped[i] {
			stack = stack[:len(stack) - 1]
			i++
		}
	}
	return !(len(stack) > 0)
}