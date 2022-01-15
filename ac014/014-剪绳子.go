package ac014

/*
题目：
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1]
可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，
此时得到的最大乘积是18。

1. 通过题目找剪裁规律
2. 动态规划，复杂度高
*/

/*
动态规划
转移方程 f(n) = max(i * max(f(n-i),n-i)),i = 1,2,...,n-1 时间复杂度O(n2)
*/
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*max(dp[i-j], i-j))
		}
	}
	return dp[n]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
找规律：基本不等式 https://www.bilibili.com/video/BV19a411A7ku?spm_id_from=333.999.0.0
根据算数均值不等式推导，尽可能按照3划分
*/

func cuttingRope1( num int ) int {
	if num == 1 || num ==2 {
		return 1
	}
	if num == 3 {
		return 2
	}
	sum := 1
	for num >4 {
		sum  = sum*3%1000000007
		num = num-3
	}
	return (sum*num)%1000000007
}

func cuttingRope2(n int) int {
	if n <= 3 {
		return n-1
	}
	// n = 3a + b
	ret := 1
	// b 等于0，1，2情况最大乘4,所以n > 4
	for n > 4 {
		ret = ret * 3 % 1000000007
		n -= 3
	}
	// 最后结果只会剩下2,3,4 所以直接乘以n再取余1000000007
	return ret * n % 1000000007
}

