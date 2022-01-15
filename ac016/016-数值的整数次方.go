package ac016


/*
1. 循环将n个x乘起来
2. 快速幂，二分思想的应用，每次指数减半，底数平方；时间复杂度o(logn) 空间复杂度o(1)
*/

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	r := 1.0
	if n<0 {
		x = 1/x
		n = -n
	}
	for n > 0 {
		if n & 1 ==1 {
			r = r * x
		}
		x *= x
		n >>= 1
	}
	return r
}