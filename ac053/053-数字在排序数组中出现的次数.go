package ac053

/*
二分法 o(logn)
*/
func search(nums []int, target int) int {
	find := func(target,left,right int) int {
		for left <= right {
			mid := (left + right)/2
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid -1
			}
		}
		return left
	}
	if len(nums) == 0 {
		return 0
	}
	l := find(target, 0, len(nums)-1)
	r := find(target + 1, l, len(nums)-1)
	return r-l
}