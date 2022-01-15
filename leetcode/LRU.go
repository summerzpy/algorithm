package leetcode

/*
底层逻辑采用两种数据结构，分别是双向链表与 map：

双向链表的作用主要是用数据在链表中的位置来表示删除的先后顺序，越靠近 head 的节点表示使用越多，越靠近尾部的表示越需要删除的，并且双向链表可以在 O(1) 的时间复杂度内完成添加、删除、移动等操作；
map 的作用是用来加速数据的查找速度，可以直接调取数据；
所以一个数据节点是存储两份，一份位于链表中，一份存储在 map 中，每个k-v数据都是以链表节点为单位出现。
1.首尾哨兵结点的使用，简化了双向链表的删除和插入操作； 2.映射里直接存储结点指针，省去了在链表中查找结点的步骤
 */
type LinkNode struct{
	key, value int
	pre, next *LinkNode
}

type LRUCache struct {
	m map[int]*LinkNode
	capacity int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{-1, -1, nil, nil}
	tail := &LinkNode{-1, -1, nil, nil}
	head.next = tail
	tail.pre = head
	cache := LRUCache{make(map[int]*LinkNode), capacity, head, tail}
	return cache
}

func (this *LRUCache) AddNode(node *LinkNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next = node
	node.next.pre = node
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

/*
进行 Get 操作时，主要是两步：

如果 key 存在，则获取数据，并将给节点移至头部；
如果不存在，返回 -1；
*/
func (this *LRUCache) Get(key int) int {
	m := this.m
	if node, ok := m[key]; ok {
		this.MoveToHead(node)
		return node.value
	} else {
		return -1
	}
}
/*
进行 Put 操作时，需要满足三步：

当 key 存在时，更新该 key 对应的 value，并且将该 k-v 放置头部；
当 key 不存在时，如果缓存没达到容量，则在map中和链表中插入k-v；
当 key 不存在时且缓存容量已满时，则删除链表尾部的 node 以及 map 中的 k-v，然后在链表头部插入node，并在 map 中插入新的 k-v；
*/
func (this *LRUCache) Put(key int, value int)  {
	m := this.m
	if node, ok := m[key]; ok {
		node.value = value
		this.MoveToHead(node)
	} else {
		n := &LinkNode{key, value, nil, nil}
		if len(m) >= this.capacity {
			delete(m, this.tail.pre.key)
			this.RemoveNode(this.tail.pre)
		}
		m[key] = n
		this.AddNode(n)
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

