package ac057

/*
滑动窗口
*/
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	l,r := 1,1
	sum := 1
	for l <= target/2 {
		if sum > target{
			sum -= l
			l++
		} else if sum < target {
			r++
			sum +=r
		} else {
			tmp := make([]int,r-l+1)
			for i := l;i<=r;i++ {
				tmp[i-l] = i
			}
			res = append(res, tmp)
			sum -= l
			l++
		}
	}
	return res
}
