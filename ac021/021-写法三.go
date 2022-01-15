package ac021

/*
快慢双指针
*/

func exchange2(nums []int) []int {
	var i, j = 0, 0
	for j < len(nums) {
		if nums[j]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return nums
}
