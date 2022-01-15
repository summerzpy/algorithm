package ac012

//标记矩阵
func hasPath2(board [][]byte, word string) bool {
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			mark:=make([][]bool,len(board))
			for k:=0;k<len(board);k++{
				mark[k]= make([]bool,len(board[0]))
			}
			if dfs2(board,mark,i,j,word){
				return true
			}
		}
	}
	return false
}

func dfs2(matrix [][]byte,mark [][]bool, i int,j int, word string) bool {

	if len(word) == 0 {
		return true
	}
	if i<len(matrix) && i>=0 && j<len(matrix[0]) && j>=0{
		if mark[i][j] == false && matrix[i][j] == byte(word[0]){
			mark[i][j] = true
			down := dfs2(matrix, mark, i+1, j ,word[1:])
			up := dfs2(matrix, mark, i-1, j ,word[1:])
			right := dfs2(matrix, mark, i, j+1 ,word[1:])
			left := dfs2(matrix, mark, i, j-1 ,word[1:])
			if down || up || right || left {
				return true
			} else {
				mark[i][j] = false
			}
		}
	}
	return false
}
