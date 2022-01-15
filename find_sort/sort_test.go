package find_sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	input := []int{3,1,7,5,2,9,2}
	quickSort(input)
	fmt.Println(input)
}

