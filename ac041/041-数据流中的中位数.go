package ac041

import "container/heap"


/*
最大堆：存放最小的k个数；最小堆：存放最大的k个数
len相等，放入minHeap；否则放入maxHeap

需要定义堆类型，一般为一个数组，然后为它实现五个简单的方法，即实现heap.Interface接口。
heap包会利用我们实现的这五个简单方法实现Push、Pop、Remove、Fix函数，之后就可以使用heap包中的Push和Pop函数了。
参考：https://blog.csdn.net/cheyo809775692/article/details/100924533

"container/heap"，在该package里面定义了Heap（堆）这一数据结构的使用接口。
只要自定义的数据类型实现了标准接口，可以很方便的对自定义的数据类型在堆中进行排序了。
需要定义五个接口函数：Push(x),Pop(),Len(),Less(i,j),Swap(i,j) Peek()
*less: 索引为 i 的元素是否应该排在索引为 j 的元素之前
*/

type MedianFinder struct {
	big *maxHeap
	small *minHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	big:=&maxHeap{}
	small:=&minHeap{}
	return MedianFinder{big,small}
}


func (this *MedianFinder) AddNum(num int)  {
	if (this.big.Len() + this.small.Len())%2 == 0 {
		//最小堆加入元素前，经最大堆选出适合元素
		heap.Push(this.big, num)
		heap.Push(this.small, heap.Pop(this.big))
	}  else {
		heap.Push(this.small, num)
		heap.Push(this.big, heap.Pop(this.small))
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.big.Len() == this.small.Len() {
		return (float64(this.big.Peek()) + float64(this.small.Peek()))/2
	} else {
		return float64(this.small.Peek())
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type maxHeap []int

func (h *maxHeap) Len() int {
	return len(*h)
}

func (h *maxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *maxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 查看堆顶元素
func (h *maxHeap) Peek() int {
	return (*h)[0]
}

type minHeap []int

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 查看堆顶元素
func (h *minHeap) Peek() int {
	return (*h)[0]
}

//java版
//import java.util.*;
//public class Solution {
//
//private PriorityQueue<Integer> left = new PriorityQueue<>((o1,o2)->o2-o1);
//private PriorityQueue<Integer> right = new PriorityQueue<>();
//private int N = 0;
//
//public void Insert(Integer num) {
//if (N % 2 == 0) {
//left.add(num);
//right.add(left.poll());
//} else {
//right.add(num);
//left.add(right.poll());
//}
//N++;
//}
//
//public Double GetMedian() {
//if (N % 2 == 0)
//return (left.peek() + right.peek()) / 2.0;
//else
//return (double) right.peek();
//}




