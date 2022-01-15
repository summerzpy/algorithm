package ac058

//双指针：去除多余空格；整体翻转；翻转每个单词
//时间复杂度：O(N);空间复杂度：O(N)

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	i := len(s) -1
	res := make([]string,0)
	for i >= 0 {
		if s[i] != ' '{
			j := i
			for i >= 0 && s[i] != ' '{
				i--
			}
			res = append(res, s[i+1:j+1])
		}
		i--
	}
	str := ""
	for k := 0;k<len(res);k++ {
		str = str + res[k] + " "
	}
	if str != ""{
		str = str[:len(str)-1]
	}
	return str
}
