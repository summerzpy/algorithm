package ac040

import "container/heap"

/*
构建大根堆；快排，取前k个；基于快排的快速选择
*/

func getLeastNumbers(arr []int, k int) []int {
	if len(arr)==0 || k==0 {
		return nil
	}

	//方法一 基于快排的快速选择
	// return quicksearch(arr, 0, len(arr)-1, k)

	//方法二 快排 然后取前k个
	// quicksort(arr, 0, len(arr)-1)
	// return arr[:k]


	//方法三  建堆，大根堆
	d := &heapInt{}
	for _,v := range arr {
		heap.Push(d, v)
		if d.Len() > k{
			heap.Pop(d)
		}
	}
	return *d


}

func partition(nums []int,i,j int) int {
	l,m,r:=i,i,j
	for l<r {
		for l<r && nums[r]>=nums[m] {
			r--
		}
		for l<r && nums[l]<=nums[m] {
			l++
		}
		if l<r {
			nums[l],nums[r]=nums[r],nums[l]
		}
	}
	nums[m],nums[l]=nums[l],nums[m]

	return l
}

func quicksearch(nums []int,i,j,k int) []int{

	t:=partition(nums,i,j)

	if t==k-1 {
		return nums[:k]
	}

	if t<k-1 {
		return quicksearch(nums,t+1,j,k)
	}


	return quicksearch(nums,i,t-1,k)

}

func quicksort(nums []int,i,j int) {
	if i>=j {
		return
	}

	m:=partition(nums, i,j)
	quicksort(nums,i,m-1)
	quicksort(nums,m+1,j)
}

type heapInt []int

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt)Less(i,j int)bool {return (*h)[i]>(*h)[j]}
func (h *heapInt)Swap(i,j int) {(*h)[i],(*h)[j]=(*h)[j],(*h)[i]}
func (h *heapInt)Len() int {return len(*h)}
func (h *heapInt)Push(x interface{}) {
	*h=append(*h,x.(int))
}
func (h *heapInt)Pop() interface{} {
	t:=(*h)[len(*h)-1]
	*h=(*h)[:len(*h)-1]
	return t
}
func (h *heapInt)Peek() int {
	return (*h)[0]
}
