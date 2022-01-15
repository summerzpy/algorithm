package ac067

import (
	"math"
	"strings"
)

func strToInt(str string) int {
	//strings.TrimSpace()返回一个string类型的slice，并将最前面和最后面的ASCII定义的空格去掉，中间的空格不会去掉
	str = strings.TrimSpace(str)
	res := 0
	sign := 1
	for i, v := range str {
		if v >= '0' && v <= '9' {
			res = res*10 + int(v-'0')
		} else if i == 0 && v == '+' {
			sign = 1
		} else if i == 0 && v == '-' {
			sign = -1
		} else {
			break
		}
		if res > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return sign * res
}
