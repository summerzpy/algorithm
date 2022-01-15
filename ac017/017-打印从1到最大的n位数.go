package ac017

/*
题目：输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。
比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
*/
func printNumbers(n int) []int {
	r := make([]int, 0)
	max := 0
	for n > 0 {
		max = max*10 + 9
		n--
	}
	for i := 1; i<=max; i++ {
		r = append(r,i)
	}
	return r
}
