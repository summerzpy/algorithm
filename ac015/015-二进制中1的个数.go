package ac015

/*
1. 逐位判断
2. 巧用n&(n-1)

位移
>>> 无符号右移，无论正负数，前面补零
>> 右移，正数前面补零，负数前面补1
<< 左移，无论正负数，后面补零
*/
func hammingWeight(num uint32) int {
	cnt := 0
	for num > 0 {
		num = num&(num-1)
		cnt++
	}
	return cnt
}

func hammingWeight1(num uint32) int {
	res := 0
	for num > 0 {
		if num&1 >0 {
			res++
		}
		num >>=1
	}
	return res
}

