package find_sort

import (
	"math/rand"
)

func sortArray2(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	randomEle := rand.Intn(right-left+1) + left
	nums[left], nums[randomEle] = nums[randomEle], nums[left]
	i, j, base := left, right, nums[left]
	for i < j {
		for nums[j] >= base && i < j {
			j--
		}
		for nums[i] <= base && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort(nums, left, i - 1)
	quickSort(nums, i + 1, right)
}

