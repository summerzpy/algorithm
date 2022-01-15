package leetcode

import "math/rand"

type RandomizedSet struct {
	dict map[int]int
	list []int
}

func Constructor1() RandomizedSet {
	dict := make(map[int]int)
	var list []int
	return RandomizedSet{dict, list}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, has := this.dict[val]; has {
		return false
	}
	this.dict[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, has := this.dict[val]; !has {
		return false
	} else {
		lastIdx := len(this.list) - 1
		lastVal := this.list[lastIdx]
		this.list[idx] = this.list[lastIdx]
		this.dict[lastVal] = idx
		this.list = this.list[:lastIdx]
		delete(this.dict, val)
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}
