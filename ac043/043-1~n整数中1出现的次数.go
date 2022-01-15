package ac043

/*
找数学规律，当前位cur是否大于1，等于1，小于1

*/

func countDigitOne(n int) int {
	//从个位开始
	base := 1
	res := 0
	//当base 大于 n，跳出循环
	for n >= base {
		a := n/base
		b := n % base
		cur := a % 10
		a = a/10
		if cur >1 {
			res = res + (a+1) * base
		} else if cur == 1 {
			res = res + a*base + (b+1)
		} else {
			res = res + a*base
		}
		//base左移一位
		base = base * 10
	}
	return res
}