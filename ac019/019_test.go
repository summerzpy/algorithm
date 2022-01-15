package ac019

import (
	"fmt"
	"testing"
)

func Test_isMatch(t *testing.T) {
	p := "A*B.A*B"
	s := "AAABCAAB"
	fmt.Println(isMatch(s, p))
}
