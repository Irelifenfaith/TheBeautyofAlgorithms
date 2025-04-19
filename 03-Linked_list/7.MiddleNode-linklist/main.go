package main

import "fmt"

// 链表中间节点

//链表节点定义
type ListNode struct {
	Value byte
	Next  *ListNode
}

//链表
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
		Value: b,
		Next:  nil,
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
		p = p.Next
	}

	// 如果链表长度大于链表容量，报错返回
	if s.length > s.capacity {
		panic("error")
	}

	// 插入目标节点至链表，其前驱节点为链表尾节点pr
	q = &ListNode{
		Value: b,
		Next:  nil,
	}
	// 插入链表尾部
	pr.Next = q
	// 更新尾节点
	s.length++
}

func middleNode(head *ListNode) *ListNode {

	// 定义快慢指针
	var fast *ListNode
	var slow *ListNode
	fast, slow = head, head

	for fast.Next != nil && fast.Next.Next != nil {
		// 快指针走两步，慢指针走1步
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 链表遍历完成后，慢指针指向的即为中间节点
	return slow
}

// 链表中间节点
func main() {
	strList := NewStrList(6)

	strList.put('s')
	strList.put('y')
	strList.put('j')
	strList.put('w')
	strList.put('p')
	strList.put('q')

	var node *ListNode = strList.head
	for ; node != nil; node = node.Next {
		fmt.Printf("%c -> ", node.Value)
	}

	fmt.Printf("List has cycle: %c", middleNode(strList.head).Value)

}
