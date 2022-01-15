package find_sort

/*
二分查找
时间复杂度：O(\log n)O(logn)，其中 nn 是数组的长度。
空间复杂度：O(1)O(1)。
*/
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

