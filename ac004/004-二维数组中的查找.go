package ac004

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	row, col := len(matrix),len(matrix[0])
	i,j := 0,col-1
	for i<row-1 && j>0 {
		if matrix[i][j] < target{
			i++
		}else if matrix[i][j] > target {
			j--
		}else {
			return true
		}
	}
	return false
}