package ac051

//无序->有序,所有需要交换的逆序对的总和，就是题目所求
//冒泡 归并；归并：将两个有序的数组合并成一个有序的数组
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param data int整型一维数组
 * @return int整型
 */
func ReversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	tmp := make([]int, len(nums))
	return merge(nums, tmp, 0, len(nums)-1) %1000000007
}

func merge(nums, tmp []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := (start + end) / 2
	leftCnt := merge(nums, tmp, start, mid)
	rightCnt := merge(nums, tmp, mid+1, end)
	mergeCnt := 0
	index := start
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] > nums[j] {
			mergeCnt = mergeCnt + mid - i + 1
			tmp[index] = nums[j]
			j++
		} else if nums[i] <= nums[j] {
			tmp[index] = nums[i]
			i++
		}
		index++
	}
	for i <= mid {
		tmp[index] = nums[i]
		i++
		index++
	}
	for j <= end {
		tmp[index] = nums[j]
		j++
		index++
	}
	for m := start; m <= end; m++ {
		nums[m] = tmp[m]
	}
	return leftCnt + rightCnt + mergeCnt

}

