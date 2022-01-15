package leetcode

import "container/list"

//因为769是质数，便于键键值分散的更好，题意说操作10^4次，所以取一个1000左右的质数都可以
//数组+链表 链地址法
const base = 769

type entry struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func Constructor2() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (m *MyHashMap) hash(key int) int {
	return key % base
}

func (m *MyHashMap) Put(key, value int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			//修改key对应的值
			e.Value = entry{key, value}
			return
		}
	}
	m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			m.data[h].Remove(e)
		}
	}
}
