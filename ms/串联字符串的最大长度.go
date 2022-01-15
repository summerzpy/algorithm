package ms

func maxLength(arr []string) int {
	maxLenStr := dfs("", arr, 0, len(arr)-1)
	return len(maxLenStr)
}

// 检查字符串中是否有重复字符
func checkUniq(str string) bool {
	if len(str) == 0 {
		return true
	}

	countMap := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		if countMap[byte(str[i])] > 0 {
			return false
		}
		countMap[byte(str[i])]++
	}

	return true
}

// dfs: str为组合后的字符串，arr为输入的字符串数组，start为当前正在处理的字符串索引，end为字符串数组最后一个位置的索引
func dfs(str string, arr []string, start, end int) string {
	// 减枝操作：根据字符串中都是小写字母的特性，如果字符串长度大于26，说明一定有重复字符
	if len(str) > 26 {
		return ""
	}
	// 递归结束条件：处理完字符串数组最后一个字符串之后
	if start > end {
		// 此时的str就是字符串就是最终合成的字符串
		// 检查其是否符合无重复字符的要求
		if checkUniq(str) {
			return str
		}
		return ""
	}

	// 当处理start位置的字符串时我们有两种选择，一种是不把arr[start]加入到合成的字符串中
	// 另外一种就是把arr[start]加入到合成的字符串中，然后继续进行dfs搜索，找到符合不重复条件
	// 并且长度最长的合成字符串的组合
	noAddStr := dfs(str, arr, start+1, end)
	addStr := ""
	// 这里进行一个减枝的操作，加入当前str+arr[start]合成的字符串中已经有重复字符了，那就没必要
	// 继续调用dfs搜索了
	if checkUniq(str + arr[start]) {
		addStr = dfs(str+arr[start], arr, start+1, end)
	}

	// 返回合成字符串最长的那个
	if len(noAddStr) > len(addStr) {
		return noAddStr
	}
	return addStr
}

