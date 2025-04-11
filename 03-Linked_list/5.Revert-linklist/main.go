package main

import "fmt"

// 反转单链表

//链表节点定义
type ListNode struct {
	value byte
	next  *ListNode
}

//双向链表
type ByteList struct {
	capacity int
	length   int
	head     *ListNode
}

// 初始化链表
func NewStrList(capacity int) *ByteList {
	return &ByteList{
		capacity: capacity,
		length:   0,
		head:     nil,
	}
}

func (s *ByteList) put(b byte) {
	//维护一个单向链表，链表中加入元素，每次在末尾插入元素
	p := s.head
	q := &ListNode{
		value: b,
		next:  nil,
	}
	var pr *ListNode
	if s.head == nil {
		// 链表中无节点
		s.head = q
		s.length++
		return
	}

	// 链表有节点，进行遍历
	for p != nil {
		pr = p
		p = p.next
	}

	// 如果链表长度大于链表容量，报错返回
	if s.length > s.capacity {
		panic("error")
	}

	// 插入目标节点至链表，其前驱节点为链表尾节点pr
	q = &ListNode{
		value: b,
		next:  nil,
	}
	// 插入链表尾部
	pr.next = q
	// 更新尾节点
	s.length++
}

func revertList(strList *ByteList) {

	var head *ListNode = strList.head
	// 当前第一个节点
	var pr *ListNode
	var next *ListNode
	// 当前节点，初始为head
	cur := head

	// 空链表或只有一个节点，直接返回即可
	if head == nil || head.next == nil {
		return
	}

	for cur != nil {
		// 保存当前节点的后继节点
		next = cur.next
		// 将当前节点的后继节点指向前驱节点
		cur.next = pr
		// 移动pr到当前节点
		pr = cur
		// 移动cur到后继节点
		cur = next
		// fmt.Printf("cur.value: %c, pr.value: %c\n", cur.value, pr.value)
	}
	strList.head = pr
}

// LRU Cache
func main() {
	strList := NewStrList(7)

	strList.put('s')
	// strList.put('y')
	// strList.put('j')
	// strList.put('w')
	// strList.put('p')

	for node := strList.head; node != nil; node = node.next {
		fmt.Printf("%c -> ", node.value)
	}
	fmt.Println("")

	revertList(strList)

	for node := strList.head; node != nil; node = node.next {
		fmt.Printf("%c -> ", node.value)
	}
	fmt.Println("")
}
