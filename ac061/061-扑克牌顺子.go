package ac061

import (
	"jz_offer/util"
	"sort"
)

/*
题目要求输入的扑克牌是顺子，则有以下几种情况：

除了大小王皆为0表示，如果有其他牌（其他数字）重复，肯定不是顺子。
如果不存在大小王0
又因扑克牌为5张，数字排序后是一个等差序列，差值为1，则最大值最小值之差必定为4
如果存在大小王0，
则最大值最小值之差不大于4 max-min<=4 即为顺子

需要的变量：最大值 最小值 大小王的数量
解题方法：排序法 非排序法
*/
func IsContinuous(nums []int) bool {
	time := make(map[int]bool)
	max, min, zero := 0, 14, 0
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			zero++
			continue
		}
		if time[nums[i]] {
			return false
		}
		max = util.Max(max, nums[i])
		min = util.Min(min, nums[i])
		time[nums[i]] = true
	}
	if zero == 0 {
		return (max - min) == 4
	}
	return (max - min) <= 4
}

//排序法
func isStraight(nums []int) bool {
	sort.Ints(nums)
	//记录每个数字之间差值的和
	sub := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i] == nums[i+1]{
			return false
		}
		sub +=nums[i+1]-nums[i]
	}
	return sub<5
}
