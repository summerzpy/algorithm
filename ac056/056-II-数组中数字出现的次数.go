package ac056

//无符号右移>>>是高位补0，移多少位补多少个0
////Go语言中，int为32位以上 因此记录64位

func singleNumber(nums []int) int {
	res:=0
	for i:=0;i<64;i++{
		count:=0
		bit:=1<<i
		for _,num:=range nums{
			if num&bit!=0{
				count++
			}
		}
		if count%3!=0{
			res|=bit
		}
	}
	return res
}
