package ac010

/*
递归
记忆化递归法
动态规划***最优***
*/

//需要额外的O(n)空间
func fib(n int) int {
	if n<=1 {
		return n
	}
	f:=make([]int,n+1)
	f[0]=0
	f[1]=1
	for i:=2;i<=n;i++{
		f[i]=(f[i-1]+f[i-2])%1000000007
	}
	return f[n]
}

func fib1(n int) int {
	a, b := 0, 1
	for ; n > 0; n-- {
		//注意这儿需要防止越界
		a, b = (a+b)%1000000007, a
	}
	return a
}
//最优！动态规划，注意循环边界
func fib2(N int) int {
	f0, f1 := 0, 1
	for i := 0; i < N; i++ {
        f0,f1 = f1,(f0 + f1) % 1000000007;
	}
	return f0
}