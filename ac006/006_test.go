package ac006

import (
	"fmt"
	"testing"
)

func Test_getTranslation(t *testing.T) {
	node1 := ListNode{1, nil}
	node2 := ListNode{3, nil}
	node3 := ListNode{2, nil}
	node1.Next = &node2
	node2.Next = &node3
	fmt.Println(reversePrint(&node1))
}
