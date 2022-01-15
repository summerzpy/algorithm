package ac013

//回溯法
/*
题目：地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

终止条件
1 i，j超范围
2 数位之和 > k
3 (i,j)已被访问过
*/

func movingCount(m int, n int, k int) int {
	visited := make([][]bool,m)
	for i := 0;i < m;i++ {
		visited[i] = make([]bool,n)
	}
	return dfs(m, n, 0, 0, k, visited)
}

func dfs(m int, n int,i int,j int, k int,visited [][]bool) int {
	if i >= m || j >= n || !movingCountCheck(i, j, k) || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(m,n,i+1,j,k,visited) + dfs(m,n,i,j+1,k,visited)
}

func movingCountCheck(m int, n int, k int) bool {
	r := 0
	for m > 0 {
		r = r + m%10
		m = m/10
	}
	for n > 0 {
		r = r + n%10
		n = n/10
	}
	if r > k {
		return false
	}
	return true
}