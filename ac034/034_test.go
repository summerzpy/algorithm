package ac034

import (
	"fmt"
	"testing"
)

func Test_pathSum(t *testing.T) {
	node2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{4, nil, nil},
		Right: &TreeNode{5, nil, nil},
	}
	node3 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{6, nil, nil},
		Right: &TreeNode{7, nil, nil},
	}
	root := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}
	path := pathSum(root,7)
	fmt.Println(path)
}
