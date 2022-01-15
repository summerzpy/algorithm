package ac056

/* 只出现一次的数
* hashmap；排序搜索法：先排序 在设置快慢指针进行比较
* 栈：先排序在借助栈判断；位运算
位运算 diff & -diff 能得到 diff 位级表示中最右侧为 1 的位，令 diff = diff & -diff。将 diff 作为区分两个元素的依据
https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/solution/yi-huo-de-miao-yong-by-peterbrave/
*/

func FindNumsAppearOnce(array []int) []int {
	res := 0
	for _, v := range array {
		//首先进行异或计算，数字本身异或为0
		res ^= v
	}
	div := 1
	for div&res == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for _, v := range array {
		if v&div == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	if a > b {
		return []int{b, a}
	} else {
		return []int{a, b}
	}
}

func FindNumsAppearOnce2(array []int) []int {
	diff := 0
	for _, v := range array {
		diff ^= v
	}
	//https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/solution/yi-huo-de-miao-yong-by-peterbrave/
	diff &= -diff
	a, b := 0, 0
	for _, v := range array {
		if v&diff == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	if a>b {
		return []int{b,a}
	} else {
		return []int{a,b}
	}
}
