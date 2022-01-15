package ac012

//题目：给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//回溯法:深度优先搜索
//两种方法存储走过的路径：m*n的标记矩阵；将走过的格子记为'#'之类不可能出现的字符
func exist(board [][]byte, word string) bool {
	if len(word) == 0 || len(board) == 0 {
		return false
	}
	for i := 0;i<len(board);i++ {
		for j := 0;j<len(board[0]);j++ {
			if dfs(board, i, j, 0, word){
				return true
			}
		}
	}
	return false
}
//注意判断条件的先后顺序
func dfs(board [][]byte, i int,j int, k int, word string) bool {
	//遍历是否越界 || 字符是否匹配
	if i<0 || i>=len(board) || j < 0 || j>=len(board[0]) || word[k]!= board[i][j]{
		return false
	}
	if k == len(word)-1 {
		return true
	}
	//board[i][j]已经匹配当前位字符，记录一下，后续恢复
	tmp := board[i][j]
	board[i][j] = '#'
	res := dfs(board,i+1,j,k+1,word)||dfs(board,i-1,j,k+1,word)||dfs(board,i,j+1,k+1,word)||dfs(board,i,j-1,k+1,word)
	board[i][j] = tmp
	return res
}


