package ac010

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	a,b,c := 1,1,0
	for i := 2;i<=n;i++ {
		c = (a+b)%1000000007
		a = b
		b = c
	}
	return c
}
//注意循环边界 起始值
func numWays2(n int) int {
	a,b:=1,1
	for i:=1;i<=n;i++{
		a,b=b,(a+b)%1000000007
	}
	return a
}
