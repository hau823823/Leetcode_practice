package lruCache

type listNode struct {
	prev *listNode // 指向前一个节点和后一个节点的指针，用于实现双向链表
	next *listNode
	key  int       // 存储节点的值和键，键用于在哈希表中查找节点
	val  int
}

type LRUCache struct {
	m map[int]*listNode // 根据 key 快速定位对应的节点
    head *listNode      // 指向链表的伪头部和伪尾部节点，用于方便地添加和删除节点
    tail *listNode
    cap int             // 缓存的容量
    len int             // 当前缓存中的项目数量
}

func Constructor(capacity int) LRUCache {
	// 创建伪头部和伪尾部节点并互相连接
	// 这样在链表为空时，head.next 指向 tail
	// tail.prev 指向 head
	fakeLast:= &listNode{val: -1, prev: nil, next: nil}
	fakeHead := &listNode{val: -2, prev: nil, next: fakeLast}
	fakeLast.prev = fakeHead

    return LRUCache{
		m: make(map[int]*listNode, capacity),
		head: fakeHead,
		tail: fakeLast,
		cap: capacity,
		len: 0,
	}
}

func (cache *LRUCache) Get(key int) int {

	v, exists := cache.m[key]
	if !exists {
		return -1
	}

	// 如果键存在，但不是链表的第一个节点（head.next
	// 则通过 swapFirst 方法将节点移动到链表头部
	if v == cache.head.next {
		return v.val
	}

	cache.swapFirst(v)

    return v.val
}

func (cache *LRUCache) Put(key int, value int)  {
	// 首先尝试通过 Get 方法获取值，如果存在（即返回值不是 -1），则直接更新链表头部节点的值。
	if cache.Get(key) != -1 {
		cache.head.next.val = value
		return
	}

	// 如果不存在，检查当前缓存长度是否达到上限
	// 如果达到上限，删除链表尾部的节点（最近最少使用的节点），并在哈希表中移除对应的键
	var cur *listNode
	if cache.len >= cache.cap {
		cache.len--
		cur = cache.tail.prev
		cur.prev.next = cache.tail
		cache.tail.prev = cur.prev
		delete(cache.m, cur.key)
	} else {
		cur = &listNode{}
	}

	// 如果未达到上限，创建一个新的节点
	// 将新节点或更新后的节点插入到链表头部。
	// 在哈希表中添加或更新键的映射。
	cache.len++
	next := cache.head.next

	cur.val = value
	cur.key = key
	cur.next = next
	cur.prev = cache.head
	cache.head.next = cur
	next.prev = cur
	cache.m[key] = cur
}

func (cache *LRUCache) swapFirst(new *listNode) {
	frst := cache.head.next

	next := new.next
	prev := new.prev

	cache.head.next = new
	new.prev = cache.head

	prev.next = next
	next.prev = prev

	frst.prev = new
	new.next = frst

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */