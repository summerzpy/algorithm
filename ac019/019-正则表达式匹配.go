package ac019

/*
题目：
请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'
表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。

动态规划
转移方程模拟： https://www.bilibili.com/video/BV1th411o7Rg?from=search&seid=302877586367505366&spm_id_from=333.337.0.0
设置二维矩阵标记状态int[n+1][m+1],包含空字符串的情况
*/


func isMatch(s string, p string) bool {
	l1, l2 := len(s), len(p)
	//对于空串处理方便
	f := make([][]bool, l1+1)
	for i := range f {
		f[i] = make([]bool, l2+1)
	}
	//f[i][0]除f[0][0]都为false，空字符串和任意不为空的字符串都不匹配
	f[0][0] = true
	//标记s串的字符位
	for i := 0; i <= l1; i++ {
		//标记p串的字符位p[j-1]和f[i][j]匹配
		for j := 1; j <= l2; j++ {
			//p串第一个不可能为*，
			if p[j-1] == '*' {
				//'*'表示它前面的字符可以出现任意次（含0次）,若o往前两个字符已经匹配当前s，则此种情况匹配 eg:AA和A*
				if f[i][j-2] {
					f[i][j] = true
				//判断p的前序串是否和当前s串匹配,即是否可以延续前序串的匹配，eg:A*和A，AAABCA和A*B.A*(确认AAABC和A*B.A*匹配后，需要判断AAABCA和A*B.A是否匹配)
				} else if i != 0 && f[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.') {
					f[i][j] = true
				}
			//前序字符串已经匹配，看当前字符是否匹配
			} else if i != 0 && f[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.') {
				f[i][j] = true
			}
		}
	}
	return f[l1][l2]
}
