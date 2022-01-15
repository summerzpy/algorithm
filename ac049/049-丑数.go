package ac049

func nthUglyNumber(n int) int {
	if n < 7 {
		return n
	}
	i2, i3, i5 := 0, 0, 0
	r := make([]int, n)
	r[0] = 1
	min := func(a,b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 1;i<n;i++ {
		v := min(r[i2]*2, min(r[i3]*3, r[i5]*5))
		r[i] = v
		if v == r[i2]*2 {
			i2++
		}
		if v == r[i3]*3 {
			i3++
		}
		if v == r[i5]*5 {
			i5++
		}
	}
	return r[len(r)-1]
}
