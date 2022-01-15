package ac037

import (
	"fmt"
	"testing"
)

func Test_Serialize(t *testing.T) {
	node2 := &TreeNode{
		Val: 2,
		Left:  &TreeNode{5, nil, nil},
		Right: &TreeNode{6, nil, nil},
	}
	node3 := &TreeNode{
		Val: 3,
		Left:  &TreeNode{7, nil, nil},
		Right: &TreeNode{8, nil, nil},
	}
	root := &TreeNode{
		Val: 1,
		Left:  node2,
		Right: node3,
	}
	//非层次遍历：1,2,5,#,#,6,#,#,3,7,#,#,8,#,#,
	str := Serialize(root)
	fmt.Println(Deserialize(str))
}
