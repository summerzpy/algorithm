package ac050

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @return int整型
// 通过ASCII码进行大小写的转化
	// 65-90（A-Z），97-122（a-z）
借助字典存放
*/
func FirstNotRepeatingChar( s string ) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
