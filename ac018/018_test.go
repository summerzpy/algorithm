package ac018

import (
	"fmt"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	node1 := ListNode{-3, nil}
	node2 := ListNode{5, nil}
	node3 := ListNode{-99, nil}
	node1.Next = &node2
	node2.Next = &node3
	fmt.Println(deleteNode(&node1, -99))
}
