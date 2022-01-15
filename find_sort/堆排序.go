package find_sort

/*
      0
   1       2
 3   4   5   6
7 8 9 10 。。。
下表为i的节点父节点下标为 (i-1)/2
下标为i的节点的左孩子下标：i * 2 +1
下标为i的节点的右孩子下标  i*2 +2
*/
func sortArray3(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {
	end := len(nums) - 1
	buildMaxHeap(nums, end)
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		maxHeapify(nums, 0, end)
	}
}

func buildMaxHeap(nums []int, len int) {
	for i := len / 2; i >= 0; i-- {
		maxHeapify(nums, i, len)
	}
}

func maxHeapify(nums []int, cur int, end int) {
	if cur > end {
		return
	}
	index := cur
	leftChild, rightChild := index*2+1, index*2+2
	if leftChild <= end && nums[leftChild] > nums[index] {
		leftChild, index = index, leftChild
	}
	if rightChild <= end && nums[rightChild] > nums[index] {
		rightChild, index = index, rightChild
	}
	if index != cur {
		nums[cur], nums[index] = nums[index], nums[cur]
		maxHeapify(nums, index, end)
	}
}

