package main

import (
	"fmt"
	"runtime"
)


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func main() {
	//input := "pwwkew"
	//fmt.Println(ac048.LengthOfLongestSubstring(input))
	testyouling()
}
//幽灵变量
func testyouling(){
	x:=1
	fmt.Println(x)
	{
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
	runtime.Gosched()
	test := map[string]TreeNode{}
	test["123"] = TreeNode{1,nil,nil}

}
