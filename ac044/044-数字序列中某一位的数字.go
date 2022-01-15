package ac044

import "strconv"

/*
找数学规律
*/

func findNthDigit(n int) int {
	digit, start := 1, 1
	for n > 9*digit*start {
		n -= 9*digit*start
		start *= 10
		digit++
	}
	num := start + (n-1)/digit
	index := (n-1)%digit
	return getDigit(num, index)
}

func getDigit(num int,index int) int {
	strNum := strconv.Itoa(num)
	r,_ := strconv.Atoi(string(strNum[index]))
	return r
}