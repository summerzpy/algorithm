package find_sort

//时间复杂度：O(N^2)，这里 N 是数组的长度；
//空间复杂度：O(1)，使用到常数个临时变量

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}