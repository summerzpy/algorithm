package ac021

/*
首尾双指针 时间复杂度O(N),空间复杂度O(1)
*/

func exchange1(nums []int) []int {
	var i, j = 0, len(nums)-1
	for i<j {
		//左指针指向要移动的偶数，是奇数(1&1=1,其余为0)i++
		if nums[i] & 1 == 1{
			i++
			continue
		}
		//右指针指向要移动的奇数
		if nums[j] & 1 == 0 {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
