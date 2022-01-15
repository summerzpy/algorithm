package ac048

//动态规划+hashmap
//滑动窗口+hashmap
//左右变量记录窗口边界，map记录字符是否在窗口中出现过
func LengthOfLongestSubstring(s string) int {
	maxLen := 0
	m := make(map[byte]bool,0)
	for left,right := 0,0; right < len(s); {
		if _,ok := m[s[right]];ok {
			delete(m, s[left])
			left++
		}else{
			m[s[right]] = true
			right++
		}
		if right - left > maxLen {
			maxLen =  right - left
		}
	}
	return maxLen
}

