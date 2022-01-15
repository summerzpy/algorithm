package ac066

func multiply(a []int) []int {
	n := len(a)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n)
	res[0] = 1
	tmp := 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * a[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		tmp *= a[i+1]
		res[i] *= tmp
	}
	return res
}
