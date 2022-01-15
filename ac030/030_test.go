package ac030

import (
	"testing"
)

func Test_MinStack(t *testing.T) {
	var minStack MinStack1
	minStack.Push1(-2)
	minStack.Push1(0)
	minStack.Push1(-3)
	minStack.Min1()
	minStack.Pop1()
	minStack.Top1()
	minStack.Min1()
}
