package ac045

import (
	"fmt"
	"sort"
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param numbers int整型一维数组
 * @return string字符串
 */

/*
需要注意：
一、证明规则是有效的
自反性：aa = aa => a=a
对称性: a>b => b<a
传递性: a<b, b<c => a<c
二、证明得到的数确实是最小的
反证法
*/
func PrintMinNumber( nums []int ) string {
	//注意写法
	sort.Slice(nums, func(i, j int) bool {
		return compareNumber(nums[i],nums[j])
	})
	res := ""
	for i:=0;i<len(nums);i++{
		res = res + strconv.Itoa(nums[i])
	}
	return res
}

func compareNumber(a,b int)bool{
	str1 := fmt.Sprintf("%d%d",a,b)
	str2 := fmt.Sprintf("%d%d",b,a)
	if str1<str2{
		return true
	}
	return false
}
