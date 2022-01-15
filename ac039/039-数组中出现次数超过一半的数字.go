package ac039

/*
哈希表：时间复杂度O(N),时间复杂度O(N)
摩尔投票：时间复杂度O(N),空间复杂度O(1)
*/

func majorityElement(nums []int) int {
	rating := 0
	res := nums[0]
	for _,v := range nums {
		if rating == 0 {
			res = v
		}
		if res == v {
			rating += 1
		}else {
			rating--
		}
	}
	return res
}