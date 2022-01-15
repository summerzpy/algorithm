package ac011

//二分法
//[3,4,5,1,2] 注意执行步骤
func minArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	l,h := 0, len(nums)-1
	for l<h {
		m := (l + h) / 2
		if nums[m] < nums[h] {
			h = m
		} else if nums[m] > nums[h] {
			l = m + 1
		} else {
			h--
		}
	}
	return nums[l]
}


