package ac056

/* 只出现一次的数
* hashmap；排序搜索法：先排序 在设置快慢指针进行比较
* 栈：先排序在借助栈判断；位运算
位运算 diff & -diff 能得到 diff 位级表示中最右侧为 1 的位，令 diff = diff & -diff。将 diff 作为区分两个元素的依据
*/

func FindNumsAppearOnce(array []int) []int {
	res := 0
	for _, v := range array {
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
