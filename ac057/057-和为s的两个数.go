package ac057

func FindNumbersWithSum(array []int, sum int) []int {
	r := []int{}
	i, j := 0, len(array)-1
	for i<j {
		cur := array[i] + array[j]
		if cur == sum {
			return []int{array[i],array[j]}
		}
		if cur >sum {
			j--
		}else {
			i++
		}
	}
	return r
}
