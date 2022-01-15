package find_sort

//思想：分而治之
//时间复杂度：O(NlogN)，这里 N 是数组的长度；
//空间复杂度：O(N)，辅助数组与输入数组规模相当。
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left) + len(right))
	l, r, i := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}
	for l < len(left) {
		result[i] = left[l]
		i++
		l++
	}
	for r < len(right) {
		result[i] = right[r]
		i++
		r++
	}
	return result
}
