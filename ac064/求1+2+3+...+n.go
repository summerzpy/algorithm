package ac064

//使用递归，最重要的是指定返回条件 灵活用&&
func Sum_Solution(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}
