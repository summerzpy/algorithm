package ac060

import (
	"math"
)

//动态规划：初始状态 递推条件 边界条件
func dicesProbability(n int) []float64 {
	dp := make([][]float64, n)
	res := make([]float64, 5*n+1)
	for i := range dp {
		dp[i] = make([]float64, (i+1)*6-i)
	}
	for i := range dp[0] {
		dp[0][i] = float64(1) / float64(6)
	}
	for i := 1; i < n; i++ { //骰子数
		for j := i; j < 6*n; j++ { //点数总和
			for k := 1; k <= 6; k++ {
				if i-1 >= 0 && j-k >= 0 {
					dp[i][j] += dp[i-1][j-k]
					//dp[i][j] = dp[i-1][j-1]+dp[i-1][j-2]+……+dp[i-1][j-6]
				}
			}
		}
	}
	fenmu := math.Pow(6, float64(n))
	for i := 0; i < 5*n+1; i++ {
		res[i] = dp[n-1][i+n-1] / fenmu
	}
	return dp[n-1]
}

//更优
func dicesProbability2(n int) []float64 {
	dp := make([][]float64, n + 1 )
	for i := range dp{
		dp[i] = make([]float64, 6*n + 1)
	}
	for i := 1;i<= 6;i++ {
		dp[1][i] = float64(1)/float64(6)
	}
	for i := 2; i < n + 1 ; i ++{
		for j :=i;j<i*6 + 1;j++{
			for k := 1;k<=6;k++{
				if j-k >=0 {
					//dp[i][j] = dp[i-1][j-1]+dp[i-1][j-2]+……+dp[i-1][j-6]
					dp[i][j] += dp[i-1][j-k]/6
				} else{
					continue
				}
			}
		}
	}
	return dp[n][n:]
}

