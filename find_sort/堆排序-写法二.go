package find_sort

//时间复杂度O(nlogn) 空间复杂度O(1)
func heapSort2(nums []int) {
	n := len(nums)

	var heapify func(root, length int)

	heapify = func(root, length int) {
		//堆顶元素下标从0开始
		//左右节点的下标
		left, right := root * 2 + 1, root * 2 + 2
		//当前擂主
		largest := root
		if left < length && nums[left] > nums[largest] {
			largest = left
		}
		if right < length && nums[right] > nums[largest] {
			largest = right
		}

		//如果擂主发生了变化，就强者上位，弱者下位并向下发起挑战赛
		if root != largest {
			//强弱更替
			nums[root], nums[largest] = nums[largest], nums[root]
			//再向下发起挑战赛
			heapify(largest, length)
		}
	}

	//先建初堆
	for i := n / 2 - 1; i >= 0; i-- {
		//擂台赛发起者为i, 参加总人数为n
		heapify(i, n)
	}

	//真正的排序操作（i就代表了当前堆的大小）
	for i := n - 1; i > 0; i-- {
		//最强之人与末尾的人进行交换
		nums[0], nums[i] = nums[i], nums[0]
		//新来的第一擂主发起擂台赛，参加人数为i
		heapify(0, i)
	}
}

