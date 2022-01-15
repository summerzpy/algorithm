package ac053

//二分查找 时间复杂度O(logn)
//[1,3,3,3,3,4,5],0

func GetNumberOfK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	l := helper(nums, 0, len(nums)-1, k)
	r := helper(nums, l, len(nums)-1, k+1)
	return r - l
}
func helper(nums []int, left int, right int, target int) int {
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
