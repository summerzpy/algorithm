package find_sort

import "math/rand"

func sortArray(nums []int) []int {
	quickSort2(nums, 0, len(nums)-1)
	return nums
}

func quickSort2(nums []int, left int, right int) {
	if left >= right {
		return
	}
	randomIndex := randomPartition(nums, left, right)
	quickSort2(nums, left, randomIndex-1)
	quickSort2(nums, randomIndex+1, right)
}

//随机取基准
func randomPartition(nums []int, left int, right int) int {
	randomEle := rand.Intn(right-left+1) + left
	nums[left], nums[randomEle] = nums[randomEle], nums[left]
	return partition(nums, left, right)
}

func partition(nums []int, left int, right int) int {
	pivot, pivotIndex := nums[left], left
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[pivotIndex], nums[left] = nums[left], nums[pivotIndex]
	return left
}

