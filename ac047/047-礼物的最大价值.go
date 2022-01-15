package ac047

//动态规划
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := make([]int, len(grid[0]) + 1)
	dp[0] = 0
	max := func(a,b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0;i < len(grid); i++ {
		for j := 1;j <len(grid[0]) +1;j++ {
			dp[j] = max(dp[j-1], dp[j]) + grid[i][j-1]
	}
	}
	return dp[len(dp)-1]
}
