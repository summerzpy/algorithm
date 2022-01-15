package ac029

/*
注意边界条件判断
*/
func spiralOrder(matrix [][]int) []int {
	res :=make([]int,0)
	if len(matrix) == 0 {
		return res
	}
	r1, r2, c1, c2 := 0, len(matrix)-1, 0, len(matrix[0])-1

	for r1 <= r2 && c1 <= c2{
		for i := c1; i < c2; i++ {
			res =append(res,matrix[r1][i])
		}
		for i := r1 ; i <= r2; i++ {
			res = append(res, matrix[i][c2])
		}
		if r2>r1 {
			for i := c2-1; i>=c1; i-- {
				res = append(res, matrix[r2][i])
			}
		}
		if c2>c1 {
			for i := r2-1; i>r1 ;i-- {
				res = append(res, matrix[i][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return res
}