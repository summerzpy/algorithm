package ac020

import "strings"
/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。

数值（按顺序）可以分成以下几个部分：

若干空格
一个小数或者整数
（可选）一个'e'或'E'，后面跟着一个整数
若干空格

*/

func isNumber2(s string) bool {
	if len(s) == 0 {
		return false
	}
	//去掉首尾空格
	s = strings.Trim(s," ")
	numFlag,dotFlag,eFlag := false, false, false
	for i := 0; i<len(s);i++ {
		//判定为数字，则标记numFlag
		if s[i] >='0' && s[i] <= '9' {
			numFlag = true
			//判定为.  需要没出现过.并且没出现过e
		} else if s[i] == '.' && !dotFlag && !eFlag {
			dotFlag = true
			//判定为e，需要没出现过e，并且出过数字了
		} else if (s[i] == 'e' || s[i] == 'E') && numFlag && !eFlag {
			eFlag = true
			//为了避免123e这种请求，出现e之后就标志为false
			numFlag = false
			//判定为+-符号，只能出现在第一位或者紧接e后面
		} else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {
			//其他情况，都是非法的
		} else {
			return false
		}
	}
	return numFlag
}
