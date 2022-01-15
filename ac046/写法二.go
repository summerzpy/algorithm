package ac046

func translateNum(num int) int {
	//标记f(i+2)
	y := num % 10
	//a,b 用于存放中间值；x,y用于表示第i，i+1位上的数字
	a,b := 1, 1
	for num > 0 {
		num = num/10
		x := num % 10
		if (x*10 + y)>9 && (x*10 + y) < 26 {
			a,b = a + b,a
		} else {
			b = a
		}
		y = x
	}
	return a
}
