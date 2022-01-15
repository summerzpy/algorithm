package ac042

/*
动态规划
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	//表示以当前数字结尾的最大子串和
	curSum := 0
	for i:= 0;i<len(nums);i++ {
		if curSum <= 0 {
			curSum = nums[i]
		} else {
			curSum = curSum + nums[i]
		}
		//判断之前的子串和是否比当前子串最大和 还要大
		res = max(res, curSum)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
