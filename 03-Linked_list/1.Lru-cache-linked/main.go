package main

import "fmt"

// 链表节点结构体
type Node struct {
	Value int
	Next  *Node
}

// LRU表
type LRUCache struct {
	capacity int
	length   int
	head     *Node // 头节点（最近使用）
}

// LRU Cache初始化
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		head:     nil,
		length:   0,
	}
}

// Get方法：查找并更新LRU
func (c *LRUCache) Get(value int) bool {
	// 思路：维护一个有序链表，越靠近链表尾部的是越早访问的。
	// 当一个新数据访问时，从链表头顺序遍历链表
	// 1.如果数据已经被缓存到链表中，则将该节点插入到链表头部；
	// 2.如果数据在链表中未找到，分两种情况：
	//   a.如果缓存未满，则直接将此节点插入到链表头部；
	//   b.如果缓存已满，则删除尾节点，将数据节点插入到头节点
	// 由于无论如何都需要遍历整个链表，因此时间复杂度是O(n)
	var prev *Node
	p := c.head
	found := false

	for p != nil {
		if p.Value == value { //查到了，进一步处理
			found = true
			break
		}
		//没查到，继续遍历，并记录p
		prev = p
		p = p.Next
	}

	if found {
		// 查到了p即查到的节点（源节点）
		// 查到了，如果是头节点，无需移动直接返回(prev为nil，说明上面循环一次就找到了，prev没有被赋值)
		if prev == nil {
			return found
		}
		//非头节点
		//断开源节点（源节点的前驱prev的后继节点=源节点的后继，如2->3->1->5， 查到了3，则前驱是2，后继是1，需变为2->1->5）
		prev.Next = p.Next
		//插入到头部（源节点的后继=链表头节点变为了3->2->1->5，）
		p.Next = c.head
		//将源节点赋值为链表头节点
		c.head = p
		return found
	} else {
		// 没查到：缓存不满，则将节点插入到链表头；缓存满了，则需将链表尾节点淘汰
		// 新增节点到链表头部
		newNode := &Node{Value: value}
		//插入到头部
		newNode.Next = c.head
		//将节点赋值为链表头节点
		c.head = newNode
		//长度+1
		c.length++
		// 缓存满了，将链表尾节点删除
		if c.length > c.capacity {
			var tailPrev *Node
			q := c.head
			//遍历链表，直至链表尾部（链表尾节点即q.Next）
			for q.Next != nil {
				tailPrev = q
				q = q.Next
			}
			// 将链表尾节点的前驱节点的后继节点赋值为nil，即删除链表尾节点
			tailPrev.Next = nil
			// 长度-1
			c.length--
		}
		return false
	}

}

// LRU Cache
func main() {
	cache := NewLRUCache(4)
	fmt.Println(cache.Get(1))
	fmt.Printf("验证头节点找到：")
	fmt.Println(cache.Get(1))
	cache.Get(2)
	fmt.Printf("Get(2)后链表顺序：")
	// 验证链表结构
	for node := cache.head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	fmt.Println("")
	cache.Get(3)
	fmt.Printf("Get(3)后链表顺序：")
	// 验证链表结构
	for node := cache.head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	fmt.Println("")
	cache.Get(4)
	fmt.Printf("Get(4)后链表顺序：")
	// 验证链表结构
	for node := cache.head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	fmt.Println("")
	cache.Get(2)
	fmt.Printf("第二次Get(2)后链表顺序：")
	// 验证链表结构
	for node := cache.head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	// 当插入第五个元素时触发删除
	cache.Get(5) // 删除最旧的节点（此时应删除1）
	fmt.Printf("\nGet(5)后链表顺序：")
	for node := cache.head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}

	fmt.Println("")
	//输出结果如下：
	// false
	// 验证头节点找到：true
	// Get(2)后链表顺序：2 -> 1 ->
	// Get(3)后链表顺序：3 -> 2 -> 1 ->
	// Get(4)后链表顺序：4 -> 3 -> 2 -> 1 ->
	// 第二次Get(2)后链表顺序：2 -> 4 -> 3 -> 1 ->
	// Get(5)后链表顺序：5 -> 2 -> 4 -> 3 ->

}
