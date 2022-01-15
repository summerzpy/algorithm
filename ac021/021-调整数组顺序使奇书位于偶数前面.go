package ac021


/*
辅助空间
*/
func exchange(nums []int) []int {
	var i, j = 0, len(nums)
	var tmp = make([]int,j)
	for _, v := range nums{
		if v & 1 == 1 {
			tmp[i] = v
			i++
		}else{
			tmp[j-1] = v
			j--
		}
	}
	//此时tmp = [1,3,4,2]
	//如果需要使奇偶数保持相对位置不变，后半部分进行翻转
	count := len(nums)-i
	for f:= 0;f<count/2;f++{
		tmp[j+f],tmp[len(nums)-1-f] = tmp[len(nums)-1-f],tmp[j+f]
	}
	return tmp
}



