package ac059

import "container/list"

//大顶堆
//public ArrayList<Integer> maxInWindows(int[] num, int size) {
//	ArrayList<Integer> ret = new ArrayList<>();
//	if (size > num.length || size < 1)
//		return ret;
//	PriorityQueue<Integer> heap = new PriorityQueue<>((o1, o2) -> o2 - o1);  /* 大顶堆 */
//	for (int i = 0; i < size; i++)
//		heap.add(num[i]);
//	ret.add(heap.peek());
//	for (int i = 0, j = i + size; j < num.length; i++, j++) {            /* 维护一个大小为 size 的大顶堆 */
//		heap.remove(num[i]);
//		heap.add(num[j]);
//		ret.add(heap.peek());
//	}
//	return ret;
//}
//双端队列：队列从队头到队尾对应窗口内从最大值到尾元素的一个子序列；队列存储窗口元素的下表值，便于判断队头出队。
//时间复杂度O(n)
func maxInWindows(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}
	queue := list.New()
	res := make([]int, 0, n-k+1) //不能是make([]int, n-k+1)  [10,14,12,11],4 期望[14] 实际[0,14]
	for i := 0; i < n; i++ {
		//先判断是否需要排出队尾元素；当队尾元素小于当前元素，排出
		for queue.Len() > 0 && nums[queue.Back().Value.(int)] <= nums[i] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(i)
		//当前值大于队头元素，推出队头元素
		if queue.Front().Value.(int) < i-k+1 {
			queue.Remove(queue.Front())
		}
		if i+1 >= k {
			res = append(res, nums[queue.Front().Value.(int)])
		}
	}
	return res
}


func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	que := []int{}
	for i := range nums {
		if len(que) > 0 && que[0] <= i - k {
			que = que[1:] //元素过期了超出了窗口范围就踢出去
		}
		//新元素进来时，判断是不是可以更新单调队列

		for len(que) > 0 && nums[que[len(que)-1]] <= nums[i] {
			que = que[:len(que)-1]
		}
		que = append(que, i)
		//超出队列容量时就可以更新答案了
		if i >= k - 1{
			ans = append(ans, nums[que[0]])
		}
	}
	return ans
}
