package ac003

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
// [2,3,1,0,2,5,3]
// 坐标交换,判断异常情况后循环比较
func Duplicate(nums []int) int {
	length := len(nums)
	if length <= 0 {
		return -1
	}
	for i := 0; i < length; i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
