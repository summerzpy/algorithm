package ac029

import (
	"fmt"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	input := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	fmt.Println(spiralOrder(input))
}
