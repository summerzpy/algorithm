package ac037

import (
	"strconv"
	"strings"
)
/*
层次遍历
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialize(root *TreeNode) string {
	str := ""
	if root == nil {
		str += "#,"
		return str
	}
	str = str + strconv.Itoa(root.Val) + ","
	str += Serialize(root.Left)
	str += Serialize(root.Right)
	return str
}

var index = -1

func Deserialize(s string) *TreeNode {
	var node *TreeNode = nil
	index++
	if index >= len(s) {
		return nil
	}
	strArray := strings.Split(s, ",")
	//最后一个index截取后为空
	if strArray[index] != "#" {
		val, _ := strconv.Atoi(strArray[index])
		node = &TreeNode{
			Val:   val,
			Left:  Deserialize(s),
			Right: Deserialize(s),
		}
	}
	return node
}
