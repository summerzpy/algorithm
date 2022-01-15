package find_sort


//思想来源：插入排序的优化。在插入排序里，如果靠后的数字较小，它来到前面就得交换多次。
//「希尔排序」改进了这种做法。带间隔地使用插入排序，直到最后「间隔」为 11 的时候，
//就是标准的「插入排序」，此时数组里的元素已经「几乎有序」了；
//希尔排序的「间隔序列」其实是一个超参数，这方面有一些研究成果，
//有兴趣的朋友可以了解一下。

func shellSort(arr []int) {
	// 找到当前数组需要用到的 Knuth 序列中的最大值
	maxKnuthNumber := 1;
	for maxKnuthNumber <= len(arr) / 3 {
		maxKnuthNumber = maxKnuthNumber * 3 + 1
	}
	//创建 间隔序列（增量序列）
	for gap:=maxKnuthNumber;gap>0;gap = (gap-1)/3{
		// 从 gap 开始，按照顺序将每个元素依次向前插入自己所在的组
		for i:=gap;i<len(arr);i++{
			currentNumber := arr[i]
			j := i - gap //记录当前元素下标的上一个下标
			for j>=0 && arr[j] > currentNumber{
				arr[j+gap] = arr[j] //后移元素
				j -= gap //更新下标
			}
			arr[j+gap] = currentNumber
		}

	}
}

