package find_sort

import (
	"fmt"
	"testing"
)
//https://leetcode-cn.com/problems/sort-an-array/solution/golang-by-xilepeng-2/
func Test_quickSort(t *testing.T) {
	input := []int{3,1,7,5,2,9,4}
	sortArray3(input)
	fmt.Println(input)
}

