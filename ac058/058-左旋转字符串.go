package ac058

func LeftRotateString(str string, n int) string {
	length := len(str)
	if length < 1 {
		return ""
	}
	n = n % length
	return str[n:length]+str[0:n]
}
