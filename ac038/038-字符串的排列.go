package ac038


/*
dfs,重点在于怎么剪枝，和为什么交换
参考：https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/solution/zi-fu-chuan-de-pai-lie-dfsdi-gui-by-jinx-b8vn/
*/
//执行用时优于写法二
func permutation(s string) []string {
	res := []string{}
	bytes := []byte(s)
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(bytes)-1 {
			res = append(res, string(bytes))
		}
		dict := map[byte]bool{}
		for i := x; i < len(bytes); i++ {
			if !dict[bytes[i]] {
				if x != i {
					bytes[x], bytes[i] = bytes[i], bytes[x]
				}
				dict[bytes[x]] = true
				dfs(x + 1)
				if x != i {
					bytes[x], bytes[i] = bytes[i], bytes[x]
				}
			}
		}
	}
	dfs(0)
	return res
}


